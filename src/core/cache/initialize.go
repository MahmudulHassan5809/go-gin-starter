package cache

import (
	"sync"
)

var (
    once      sync.Once
    instance  *CacheManager
)


func SetCacheManager(backend CacheBackend) {
    once.Do(func() {
        instance = NewCacheManager(backend)
    })
}


func GetCacheManager() *CacheManager {
    return instance
}
