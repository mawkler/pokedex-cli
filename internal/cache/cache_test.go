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
	cache := NewCache(time.Millisecond * 10)
	cache.Add("key", []byte("value"))
	go cache.reapLoop() // TODO: this doesn't work yet

	assert.Equal(t, 1, len(cache.cache))

	time.Sleep(time.Millisecond * 20)

	assert.Equal(t, 0, len(cache.cache))

}
