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
