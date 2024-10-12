package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Redis struct {
	Client *redis.Client
}

func NewClient() *Redis {
	rdb := &Redis{}

	rdb.init()
	return rdb
}

func (r *Redis) init() {
	r.Client = redis.NewClient(r.GetConfig())
}

func (r *Redis) GetConfig() *redis.Options {
	return &redis.Options{
		Addr:     "192.168.152.128:6379",
		Password: "",
		DB:       1,
	}
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

// SetNX key的值如果存在，则不做任何操作
func (r *Redis) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return r.Client.SetNX(ctx, key, value, expiration).Result()
}

func (r *Redis) Del(ctx context.Context, keys ...string) error {
	return r.Client.Del(ctx, keys...).Err()
}
