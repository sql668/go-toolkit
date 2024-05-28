package gtredis

import (
	"context"
	"github.com/sql668/go-toolkit/contrib/redis/locker"
	"time"
)

// NewRedis 初始化locker
func NewRedis(c *redis.Client) *Redis {
	return &Redis{
		client: c,
	}
}

type Redis struct {
	client *redis.Client
	mutex  *locker.Client
}

func (Redis) String() string {
	return "redis"
}

func (r *Redis) Lock(key string, ttl int64, options *locker.Options) (*locker.Lock, error) {
	if r.mutex == nil {
		r.mutex = locker.New(r.client)
	}
	return r.mutex.Obtain(context.TODO(), key, time.Duration(ttl)*time.Second, options)
}
