package cache

import (
	"context"
	"log"
)


type CacheManager struct {
    backend CacheBackend
}


func NewCacheManager(backend CacheBackend) *CacheManager {
    return &CacheManager{
        backend: backend,
    }
}


func (cm *CacheManager) Get(ctx context.Context, key string) (string, error) {
    data, err := cm.backend.Get(ctx, key)
    if err != nil {
        log.Printf("Error fetching key %s: %v", key, err)
        return "", err
    }
    return data, nil
}


func (cm *CacheManager) Set(ctx context.Context, key, value string, ttl int) error {
    return cm.backend.Set(ctx, key, value, ttl)
}


func (cm *CacheManager) Delete(ctx context.Context, key string) error {
    return cm.backend.Delete(ctx, key)
}


func (cm *CacheManager) GetAndDelete(ctx context.Context, key string) (string, error) {
    return cm.backend.GetAndDelete(ctx, key)
}


func (cm *CacheManager) HGetAll(ctx context.Context, key string) (map[string]string, error) {
    return cm.backend.HGetAll(ctx, key)
}


func (cm *CacheManager) HMSet(ctx context.Context, key string, fields map[string]string, ttl int) error {
    return cm.backend.HMSet(ctx, key, fields, ttl)
}
