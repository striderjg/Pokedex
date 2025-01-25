package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data map[string]cacheEntry
	mu   sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	myCache := Cache{}
	go myCache.reapLoop(interval)
	return &myCache
}

func (c *Cache) Add(key string, value []byte) {
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		c.mu.Lock()
		for key, val := range c.data {
			if val.createdAt.Add(interval).After(time.Now()) {
				// Nuke the cache item Create+interval > now
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}
