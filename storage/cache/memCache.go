package cache

import (
	"context"
	"fmt"
	"time"
)

type memCache struct {
		cache Cache
		maxSize   int64 // 最大内存，单位byte字节
		maxSizeStr string // SetMaxMemory设置的最大内存，为冗余字段
		usedSize int64 //已使用内存大小，单位字节
		ctx       context.Context
}

func NewMemCache() Cache{
	return &memCache{}
}

func WithContext(ctx context.Context, c Cache) Cache{
	return &memCache{
		cache: c,
		ctx:ctx,
	}
}

// SetMaxMemory 设置最大内存
func(mc *memCache)SetMaxMemory(m string) error{
	fmt.Println("SetMaxMemory: ",m)
	num,err := ParseSize(m)
	if err != nil {
		return err
	}
	mc.maxSizeStr = m
	mc.maxSize = num
	return nil
}

// Get 根据key获取value
func(mc *memCache)Get(key string) (interface{},  bool){

	return nil,false
}
// Set 设置值
func(mc *memCache)Set( key string, val interface{}, d time.Duration) (bool,error){
return false,nil
}
// Exists 判断key是否存在
func(mc *memCache)Exists(key string) bool{
return false
}
// Del 删除key
func(mc *memCache)Del(key string) (interface{},error){
	return nil,nil
}

// Flush 删除所有key
func(mc *memCache)Flush() bool{
	return false
}

// Keys 获取缓存中key的数量
func(mc *memCache)Keys() int64{
	return 0
}
