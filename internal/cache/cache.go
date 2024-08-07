package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	cache    map[string]cacheEntry
	lock     sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) Cache {
	return Cache{map[string]cacheEntry{}, sync.Mutex{}, interval}
}

func (cache *Cache) Add(key string, value []byte) {
	createdAt := time.Now()
	cache.cache[key] = cacheEntry{createdAt, value}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	entry, ok := cache.cache[key]
	return entry.value, ok
}

func (cache *Cache) purge() {
	for key, entry := range cache.cache {
		if entry.createdAt.Add(cache.interval).Before(time.Now()) {
			delete(cache.cache, key)
		}
	}
}

func (cache *Cache) reapLoop() {
	go func() {
		ticker := time.NewTicker(cache.interval)
		defer ticker.Stop()

		for range ticker.C {
			cache.purge()
		}
	}()
}
