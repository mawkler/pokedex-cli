package cache

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var cache Cache

func TestMain(m *testing.M) {
	cache = NewCache(time.Millisecond * 100)
	code := m.Run()

	os.Exit(code)
}

func TestAdd(t *testing.T) {
	assert.Equal(t, len(cache.cache), 0)

	cache.Add("key", []byte("value"))

	assert.Equal(t, 1, len(cache.cache))
}

func TestGet(t *testing.T) {
	cache := NewCache(time.Millisecond * 100)

	t.Run("Gets value", func(t *testing.T) {
		cache.Add("key", []byte("value"))

		value, ok := cache.Get("key")

		assert.Equal(t, true, ok)
		assert.Equal(t, []byte("value"), value)
	})

	t.Run("Doesn't get non-existent value", func(t *testing.T) {
		value, ok := cache.Get("doesn't exist")

		assert.Equal(t, false, ok)
		assert.Equal(t, []byte(nil), value)
	})
}

func TestPurge(t *testing.T) {
	interval := 5 * time.Millisecond
	waitTime := 10 * time.Millisecond

	t.Run("Purge cache", func(t *testing.T) {
		cache := NewCache(interval)
		cache.Add("key", []byte("value"))
		assert.Equal(t, 1, len(cache.cache))

		time.Sleep(waitTime)
		cache.purge()

		assert.Equal(t, 0, len(cache.cache))
	})

	t.Run("Don't purge cache that's not expired", func(t *testing.T) {
		cache := NewCache(interval)

		cache.Add("key", []byte("value"))
		assert.Equal(t, 1, len(cache.cache))

		// Note that we don't wait for cache to expire here
		cache.purge()

		assert.Equal(t, 1, len(cache.cache))
	})

	t.Run("Purge cache with reap loop", func(t *testing.T) {
		cache := NewCache(time.Millisecond * 10)
		cache.Add("key", []byte("value"))
		cache.reapLoop()

		assert.Equal(t, 1, len(cache.cache))

		time.Sleep(waitTime)

		assert.Equal(t, 0, len(cache.cache))
	})
}
