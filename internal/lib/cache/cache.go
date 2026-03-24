package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type AppCache struct {
	client      *redis.Client
	defaultTime time.Duration
}

func NewAppCache(client *redis.Client, ttl time.Duration) *AppCache {
	return &AppCache{
		client:      client,
		defaultTime: ttl,
	}
}

func (c *AppCache) Get(ctx context.Context, key string, dest any) error {
	data, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

func (c *AppCache) Set(ctx context.Context, key string, value any, ttl ...time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	expiration := c.defaultTime
	if len(ttl) > 0 {
		expiration = ttl[0]
	}

	return c.client.Set(ctx, key, data, expiration).Err()
}
