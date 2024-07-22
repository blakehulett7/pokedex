package pokecache

import (
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

func NewCache(interval time.Duration) Cache {
	return Cache{
		Map:      map[string]cacheEntry{},
		Mutex:    &sync.Mutex{},
		Interval: interval,
	}
}

func (c Cache) Add(key string, value []byte) {
	c.Map[key] = cacheEntry{
		CreatedAt: time.Now(),
		Val:       value,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	value, exists := c.Map[key]
	if !exists {
		return nil, false
	}
	return value.Val, true
}
