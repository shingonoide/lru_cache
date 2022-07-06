package cache_lru

import "time"

type Data map[string]string
type LastTimeUsed map[string]int64

type CacheLRU interface {
	Write(key string, value string)
	Read(key string) string
	Length() int
	Remove(key string)
	Clear()
	Data() Data
}

type cacheLRU struct {
	data         Data
	lastTimeUsed LastTimeUsed
	size         int
}

func (c *cacheLRU) Write(key string, value string) {
	if len(c.data) < c.size {
		c.data[key] = value
		c.lastTimeUsed[key] = time.Now().UnixNano()
	} else {
		// expires LRU
		var lru int64
		var lruKey string
		for k, v := range c.lastTimeUsed {
			if lru == 0 || v < lru {
				lru = v
				lruKey = k
			}
		}
		delete(c.data, lruKey)
		delete(c.lastTimeUsed, lruKey)
		c.data[key] = value
	}
}

func (c cacheLRU) Read(key string) string {
	c.lastTimeUsed[key] = time.Now().UnixNano()
	return c.data[key]
}

func (c cacheLRU) Length() int {
	return len(c.data)
}

func (c cacheLRU) Remove(key string) {
	delete(c.data, key)
	delete(c.lastTimeUsed, key)
}

func (c *cacheLRU) Clear() {
	c.data = make(Data)
	c.lastTimeUsed = make(LastTimeUsed)
}

func (c cacheLRU) Data() Data {
	return c.data
}

func NewCacheLRU(size int) CacheLRU {
	return &cacheLRU{
		data:         make(Data),
		lastTimeUsed: make(LastTimeUsed),
		size:         size,
	}
}
