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
	lock     *sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{map[string]cacheEntry{}, &sync.Mutex{}, interval}
	go cache.reapLoop()
	return cache
}

func (cache *Cache) Add(key string, value []byte) {
	createdAt := time.Now()

	cache.lock.Lock()
	defer cache.lock.Unlock()

	cache.cache[key] = cacheEntry{createdAt, value}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	entry, exists := cache.cache[key]
	return entry.value, exists
}

func (cache *Cache) purge() {
	for key, entry := range cache.cache {
		if entry.createdAt.Add(cache.interval).Before(time.Now()) {
			cache.lock.Lock()
			defer cache.lock.Unlock()

			delete(cache.cache, key)
		}
	}
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()

	for range ticker.C {
		cache.purge()
	}
}
