package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data map[string]cacheEntry
	mu   *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	myCache := Cache{
		data: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}
	go myCache.reapLoop(interval)
	return &myCache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	fmt.Println("Entering reapLoop")
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		fmt.Println("reapLoop Unblocked")
		c.mu.Lock()
		for key, val := range c.data {
			fmt.Printf("Created+interval: %v --- Now:  %v", val.createdAt.Add(interval), time.Now())
			if val.createdAt.Add(interval).Before(time.Now()) {
				// Nuke the cache item Create+interval > now
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}
