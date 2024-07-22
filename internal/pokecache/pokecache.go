package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Map      map[string]cacheEntry
	Mutex    *sync.Mutex
	Interval time.Duration
}

func NewCache(intervalSeconds time.Duration) Cache {
	interval := intervalSeconds * 1000000000
	return Cache{
		Map:      map[string]cacheEntry{},
		Mutex:    &sync.Mutex{},
		Interval: interval,
	}
}

func (c Cache) Add(key string, value []byte) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Map[key] = cacheEntry{
		CreatedAt: time.Now(),
		Val:       value,
	}
	return
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	value, exists := c.Map[key]
	if !exists {
		return nil, false
	}
	return value.Val, true
}

func (c Cache) ReapLoop() {
	ticker := time.NewTicker(c.Interval)
	for {
		<-ticker.C
		fmt.Println("Purging Cache...")
		c.PurgeCache()
	}
}

func (c Cache) PurgeCache() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	fmt.Println("Cache Purged!")
	return
}
