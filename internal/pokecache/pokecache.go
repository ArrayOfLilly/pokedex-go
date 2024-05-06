package pokecache

import (
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	// mu sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}


func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		// mu: sync.RWMutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	// c.mu.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: val,
	}
	// c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// c.mu.RLock()
	cachE, ok := c.cache[key]
	// c.mu.RUnlock()
	return cachE.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}	
}

func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(timeAgo) {
			// c.mu.Lock()
			delete(c.cache, k)
			// c.mu.Unock()
		}
	} 
}
