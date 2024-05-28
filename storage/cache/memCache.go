package cache

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"strconv"
	"sync"
	"time"
)

var _ Cache = (*memCache)(nil)

type memCache struct {
		items *sync.Map
	    mutex sync.RWMutex
		clearInterval time.Duration
}

// Item  每一个缓存value 都有一个value和过期时间
type item struct {
	Value   string    //这里可以使用 github.com/spf13/cast cast.ToStringE(val) 将value转换为字符串存储
	ExpiredTime time.Time   //过期时间
	Expired  int  // 设定的缓存有效时常，单位 s，如果为0，就代表永久有效
}

func NewMemCache(timer ...time.Duration) Cache{
	var mc *memCache
	if len(timer) > 0 {
		mc = &memCache{
			items: new(sync.Map),
			clearInterval: timer[0],
		}
		go mc.clearExpiredItem()
	}else{
		mc = &memCache{
			items: new(sync.Map),
		}
	}
	return mc
}

//func WithContext(ctx context.Context, c Cache) Cache{
//	return &memCache{
//		cache: c,
//		ctx:ctx,
//	}
//}

// Get 根据key获取value,如果没有对应的key或者key已经过期返回空字符串
func(mc *memCache)Get(key string) (string,  error){
	item,err := mc.getItem(key)
	if err != nil{
		return "",err
	}
	if item == nil {
		return "",errors.New(fmt.Sprintf("key: %s not existed or expired",key))
	}
	return item.Value,nil
}
// Set 设置值
func(mc *memCache)Set( key string, val interface{}, expire int) (error){
	s, err := cast.ToStringE(val)
	if err != nil {
		return err
	}

	v := &item{
		Value: s,
		ExpiredTime: time.Now().Add(time.Duration(expire) * time.Second),
		Expired: expire,
	}
	return mc.setItem(key, v)
}
// Exists 判断key是否存在
func(mc *memCache)Exists(key string) bool{
	item,err := mc.getItem(key)
	if err != nil || item == nil {
		return false
	}
	return true
}
// Del 删除key
func(mc *memCache)Del(key string) error{
	return mc.del(key)
}

// Expire 设置过期时间
func (m *memCache) Expire(key string, expire int) error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	item, err := m.getItem(key)
	if err != nil {
		return err
	}
	if item == nil {
		err = fmt.Errorf("%s not exist", key)
		return err
	}
	item.ExpiredTime = time.Now().Add(time.Duration(expire) * time.Second)
	item.Expired = expire
	return m.setItem(key, item)
}

// Flush 删除所有key
func(mc *memCache)Flush() bool{
	// 清空sync.Map
	mc.items.Range(func(key, value interface{}) bool {
		mc.items.Delete(key)
		return true
	})
	/// 验证是否清空
	//syncMap.Range(func(key, value interface{}) bool {
	//	fmt.Println("Sync.Map is empty")
	//	return false
	//})
	//mc.usedSize = 0
	return true
}

// Keys 获取缓存中key的数量
func(mc *memCache)Keys() int{

	return 0
}

func (mc *memCache) HashSet(key string, values ...interface{}) error{
	return  nil
}
func (mc *memCache) HashGet(hk, key string) (string, error){
	return "",nil
}
func (mc *memCache) HashGetALl(key string) (map[string]string, error){
	return nil,nil
}

func (mc *memCache) HashDel(hk, key string) error{
	return nil
}
func (mc *memCache) Increase(key string) error{
	return  nil
}
func (mc *memCache) Decrease(key string) error{
	return nil
}


func (mc *memCache) getItem(key string) (*item,error) {
	var err error
	i, ok := mc.items.Load(key)
	if !ok {
		return nil, nil
	}
	switch i.(type) {
	case *item:
		item := i.(*item)
		if item.Expired > 0 && item.ExpiredTime.Before(time.Now()) {
			//过期
			mc.del(key)
			//过期后删除
			return nil, nil
		}
		return item, nil
	default:
		err = fmt.Errorf("value of %s type error", key)
		return nil, err
	}
}

func (mc *memCache) del(key string) error {
	mc.items.Delete(key)
	return nil
}


func (mc *memCache) setItem(key string,item *item) error {
	mc.items.Store(key, item)
	return nil
}

func (m *memCache) calculate(key string, num int) error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	item, err := m.getItem(key)
	if err != nil {
		return err
	}

	if item == nil {
		err = fmt.Errorf("%s not exist", key)
		return err
	}
	var n int
	n, err = cast.ToIntE(item.Value)
	if err != nil {
		return err
	}
	n += num
	item.Value = strconv.Itoa(n)
	return m.setItem(key, item)
}

func (mc *memCache) clearExpiredItem(){
	timeTicker := time.NewTicker(mc.clearInterval)
	defer timeTicker.Stop()

	for {
		select {
			case <- timeTicker.C:
				mc.items.Range(func(key, value interface{}) bool {
					k,ok := key.(string)
					if ok {
						mc.getItem(k)
					}

					return true
				})
		}
	}
}
