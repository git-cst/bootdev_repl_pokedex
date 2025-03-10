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

	newCache.StartReapLoop()

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

func (c *Cache) StartReapLoop() {
	ticker := time.NewTicker(c.Interval)
	go func() {
		defer ticker.Stop()
		for range ticker.C {
			c.Reap()
		}
	}()
}

func (c *Cache) Reap() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	now := time.Now()
	for key, entry := range c.Entries {
		if now.Sub(entry.CreatedAt) > c.Interval {
			delete(c.Entries, key)
		}
	}
}
