package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Client struct {
	Client *redis.Client
}

func NewClient() *Client {
	rdb := &Client{}

	rdb.init()
	return rdb
}

func (r *Client) init() {
	r.Client = redis.NewClient(r.GetConfig())
}

func (r *Client) GetConfig() *redis.Options {
	return &redis.Options{
		Addr:     "192.168.152.128:6379",
		Password: "",
		DB:       1,
	}
}

func (r *Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *Client) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

// SetNX key的值如果存在，则不做任何操作
func (r *Client) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return r.Client.SetNX(ctx, key, value, expiration).Result()
}

func (r *Client) Del(ctx context.Context, keys ...string) error {
	return r.Client.Del(ctx, keys...).Err()
}
