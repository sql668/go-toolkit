package cache

import (
	"context"
	"time"
)


// Cache is the interface that wraps the cache.
type Cache interface {
	// SetMaxMemory 最大内存
	SetMaxMemory(size string) error
	// Get 根据key获取value
	Get(key string) (interface{},  bool)
	// Set 设置值
	Set( key string, val interface{}, d time.Duration) (bool,error)
	// Exists 判断key是否存在
	Exists(key string) bool
	// Del 删除key,并将该元素返回
	Del(key string) (interface{},error)
	// Flush 删除所有key
	Flush() bool
	// Keys 获取缓存中key的数量
	Keys() int64
}

//type cache struct {
//	cache Cache
//	ctx       context.Context
//}


type cache struct {
	cache    Cache
	ctx       context.Context
}