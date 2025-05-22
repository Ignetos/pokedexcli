package internal

import (
	"sync"
	"time"
)

type Cache struct {
	Data     map[string]cacheEntry
	mu       *sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Data:     make(map[string]cacheEntry),
		mu:       &sync.RWMutex{},
		interval: interval,
	}
	cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, b []byte) {
	c.mu.Lock()
	c.Data[key] = cacheEntry{
		createdAt: time.Now(),
		value:     b,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	data, ok := c.Data[key]
	defer c.mu.RUnlock()
	return data.value, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	go func() {
		for {
			<-ticker.C
			c.reap(c.interval)
		}
	}()
}

func (c *Cache) reap(interval time.Duration) {
	now := time.Now()
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.Data {
		if now.Sub(v.createdAt) > interval {
			delete(c.Data, k)
		}
	}
}
