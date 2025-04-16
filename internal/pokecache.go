package internal

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Cache    map[string]cacheEntry
	interval time.Duration
	mu       sync.RWMutex
	done     chan struct{}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exists := c.Cache[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) Stop() {
	close(c.done)
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for key, entry := range c.Cache {
				if time.Since(entry.createdAt) > c.interval {
					delete(c.Cache, key)
				}
			}
			c.mu.Unlock()
		case <-c.done:
			return
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Cache:    make(map[string]cacheEntry),
		interval: interval,
		mu:       sync.RWMutex{},
		done:     make(chan struct{}),
	}
	go cache.reapLoop()
	return cache
}
