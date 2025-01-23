package cache

import "context"


type CacheBackend interface {
    Get(ctx context.Context, key string) (string, error)
    Set(ctx context.Context, key string, value string, ttl int) error
    Delete(ctx context.Context, key string) error
    HGetAll(ctx context.Context, key string) (map[string]string, error)
    HMSet(ctx context.Context, key string, fields map[string]string) error
    GetAndDelete(ctx context.Context, key string) (string, error)
}