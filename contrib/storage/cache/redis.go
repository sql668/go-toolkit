package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sql668/go-toolkit/storage/cache"
	_ "github.com/sql668/go-toolkit/storage/cache"
	"time"
)

var _ cache.Cache = (*Redis)(nil)

// NewRedis redis模式
func NewRedis(client *redis.Client, options *redis.Options) (*Redis, error) {
	if client == nil {
		client = redis.NewClient(options)
	}
	r := &Redis{
		client: client,
	}
	err := r.connect()
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Redis cache implement
type Redis struct {
	client *redis.Client
}

func (*Redis) String() string {
	return "redis"
}

// connect connect test
func (r *Redis) connect() error {
	var err error
	_, err = r.client.Ping(context.TODO()).Result()
	return err
}

// Get from key
func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(context.TODO(), key).Result()
}

// Set value with key and expire time
func (r *Redis) Set(key string, val interface{}, expire int) error {
	return r.client.Set(context.TODO(), key, val, time.Duration(expire)*time.Second).Err()
}

// Del delete key in redis
func (r *Redis) Del(key string) error {
	return r.client.Del(context.TODO(), key).Err()
}

// Flush 清空
func (r *Redis) Flush() bool {
	return false
}

// HashSet set hash
func (r *Redis) HashSet(key string, values ...interface{}) error {
	_, err := r.client.HSet(context.TODO(), key, values).Result()
	return err
}

// HashGet from key
func (r *Redis) HashGet(hk, key string) (string, error) {
	return r.client.HGet(context.TODO(), hk, key).Result()
}

// HashGet from key
func (r *Redis) HashGetALl(key string) (map[string]string, error) {
	return r.client.HGetAll(context.TODO(), key).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *Redis) HashDel(hk, key string) error {
	return r.client.HDel(context.TODO(), hk, key).Err()
}

// Increase
func (r *Redis) Increase(key string) error {
	return r.client.Incr(context.TODO(), key).Err()
}

func (r *Redis) Decrease(key string) error {
	return r.client.Decr(context.TODO(), key).Err()
}

// Set ttl
func (r *Redis) Expire(key string, expire int) error {
	var dur time.Duration
	dur = time.Duration(expire) * time.Second
	return r.client.Expire(context.TODO(), key, dur).Err()
}

// GetClient 暴露原生client
func (r *Redis) GetClient() *redis.Client {
	return r.client
}
