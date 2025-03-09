package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Entries  map[string]CacheEntry
	Interval time.Duration
	Mutex    sync.Mutex
}

func NewCache() *Cache {
	newCache := &Cache{
		Entries:  make(map[string]CacheEntry),
		Interval: time.Duration(10 * time.Second),
		Mutex:    sync.Mutex{},
	}

	go newCache.reap()

	return newCache
}

func (c *Cache) Add(key string, value []byte) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Entries[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	entry, exists := c.Entries[key]
	if !exists {
		return []byte{}, false
	}

	return entry.Val, true
}

func (c *Cache) reap() {
	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()

	for range ticker.C {
		c.Mutex.Lock()
		now := time.Now()
		for key, entry := range c.Entries {
			if now.Sub(entry.CreatedAt) > c.Interval {
				delete(c.Entries, key)
			}
		}
		c.Mutex.Unlock()
	}
}
