package services

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisCache implements the Cache interface using Redis
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache creates a new Redis cache instance
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

// Get retrieves a value from Redis cache
func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

// Set stores a value in Redis cache with expiration
func (c *RedisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

// Delete removes a value from Redis cache
func (c *RedisCache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

// Clear removes all values from Redis cache
func (c *RedisCache) Clear(ctx context.Context) error {
	return c.client.FlushAll(ctx).Err()
}
