package cache

import "testing"

func TestMemcache(t *testing.T){
	cache := NewMemCache()
	cache.SetMaxMemory("100MB")
}
