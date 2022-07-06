package cache_lru_test

import (
	"testing"

	"github.com/shingonoide/lru_cache/internal/cache_lru"
	"github.com/stretchr/testify/assert"
)

func TestCacheLRU(t *testing.T) {
	t.Run("should write a key and value", func(t *testing.T) {
		cache := cache_lru.NewCacheLRU(10)
		cache.Write("abc", "def")
		assert.Equal(t, "def", cache.Read("abc"))
	})

	t.Run("should return empty value", func(t *testing.T) {
		cache := cache_lru.NewCacheLRU(10)
		assert.Equal(t, "", cache.Read("abc"))
	})

	t.Run("should return total cache itens", func(t *testing.T) {
		cache := cache_lru.NewCacheLRU(10)
		cache.Write("abc", "def")
		cache.Write("cdb", "def")
		assert.Equal(t, 2, cache.Length())
	})

	t.Run("should replace a key", func(t *testing.T) {
		cache := cache_lru.NewCacheLRU(10)
		cache.Write("abc", "def")
		cache.Write("abc", "def2")
		assert.Equal(t, 1, cache.Length())
	})

	t.Run("should remove a key", func(t *testing.T) {
		cache := cache_lru.NewCacheLRU(10)
		cache.Write("abc", "def")
		cache.Write("cdb", "def")
		cache.Remove("abc")
		assert.Equal(t, 1, cache.Length())
	})

	t.Run("should have 3 keys", func(t *testing.T) {
		cache := cache_lru.NewCacheLRU(3)
		cache.Write("abc", "def")
		cache.Write("cdb", "def")
		cache.Write("esd", "def")
		cache.Write("zxc", "def")
		assert.Equal(t, 3, cache.Length())
	})

	t.Run("should remove a key", func(t *testing.T) {
		cache := cache_lru.NewCacheLRU(10)
		cache.Write("abc", "def")
		cache.Write("cdb", "def")
		cache.Remove("abc")
		assert.Equal(t, 1, cache.Length())
	})

	t.Run("should clear all keys", func(t *testing.T) {
		cache := cache_lru.NewCacheLRU(10)
		cache.Write("abc", "def")
		cache.Write("cdb", "def")
		cache.Clear()
		assert.Equal(t, 0, cache.Length())
	})

	t.Run("should return all data", func(t *testing.T) {
		cache := cache_lru.NewCacheLRU(10)
		cache.Write("abc", "def")
		cache.Write("cdb", "def")
		assert.Equal(t, cache_lru.Data{
			"abc": "def",
			"cdb": "def",
		}, cache.Data())
	})

	t.Run("should expired least recently used item", func(t *testing.T) {
		cache := cache_lru.NewCacheLRU(2)
		cache.Write("abc", "def")
		cache.Write("bcd", "ghi")
		cache.Write("asd", "qwe")
		assert.Equal(t, "", cache.Read("abc"))
	})

}
