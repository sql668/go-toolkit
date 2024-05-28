package cache

import "context"

// Cache is the interface that wraps the cache.
type Cache interface {
	// Get 根据key获取value
	Get(key string) (string,  error)
	// Set 设置值
	Set( key string, val interface{}, expired int) (error)
	// Del 删除key
	Del(key string) error
	// Expire 设置过期时间
	Expire(key string, expired int) error
	HashSet(key string, values ...interface{}) error
	HashGet(hk, key string) (string, error)
	HashGetALl(key string) (map[string]string, error)
	HashDel(hk, key string) error
	Increase(key string) error
	Decrease(key string) error
}

type cache struct{
	cache Cache
	ctx       context.Context
}

// Get 根据key获取value
func (c *cache) Get(key string) (string,  error) {
	return c.cache.Get(key)
}
// Set 设置值
func (c *cache) Set( key string, val interface{}, expired int) (error){
	return c.cache.Set(key,val,expired)
}

// Del 删除key
func (c *cache) Del(key string) error{
	return c.cache.Del(key)
}


func (c *cache) Expire(key string, expired int) error{
	return c.cache.Expire(key,expired)
}

func (c *cache) HashSet(key string, values ...interface{}) error{
	return c.cache.HashSet(key,values)
}
func (c *cache) HashGet(hk, key string) (string, error){
	return c.cache.HashGet(hk,key)
}
func (c *cache) HashGetALl(key string) (map[string]string, error){
	return c.cache.HashGetALl(key)
}
func (c *cache) HashDel(hk, key string) error{
	return c.cache.HashDel(hk,key)
}
func (c *cache) Increase(key string) error{
	return  c.cache.Increase(key)
}
func (c *cache) Decrease(key string) error{
	return c.cache.Decrease(key)
}

func WithContext(ctx context.Context, c Cache) Cache {
	switch v := c.(type) {
	default:
		return &cache{cache: c, ctx: ctx}
	case *cache:
		lv := *v
		lv.ctx = ctx
		return &lv
	}
}

