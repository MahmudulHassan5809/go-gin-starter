package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)



type RedisBackend struct {
    client *redis.Client
}

func NewRedisBackend(url string) CacheBackend {
    opt, err := redis.ParseURL(url)
    if err != nil {
        panic(err)
    }
    return &RedisBackend{
        client: redis.NewClient(opt),
    }
}

func (r *RedisBackend) Get(ctx context.Context, key string) (string, error) {
    return r.client.Get(ctx, key).Result()
}

func (r *RedisBackend) Set(ctx context.Context, key string, value string, ttl int) error {
    return r.client.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()
}

func (r *RedisBackend) Delete(ctx context.Context, key string) error {
    return r.client.Del(ctx, key).Err()
}

func (r *RedisBackend) HGetAll(ctx context.Context, key string) (map[string]string, error) {
    return r.client.HGetAll(ctx, key).Result()
}

func (r *RedisBackend) HMSet(ctx context.Context, key string, fields map[string]string) error {
    return r.client.HSet(ctx, key, fields).Err()
}

func (r *RedisBackend) GetAndDelete(ctx context.Context, key string) (string, error) {
    val, err := r.client.Get(ctx, key).Result()
    if err != nil {
        return "", err
    }
    if err := r.client.Del(ctx, key).Err(); err != nil {
        return "", err
    }
    return val, nil
}