package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type cache struct {
	Map   map[string]cacheEntry
	Mutex *sync.Mutex
}

func NewCache(interval time.Duration) {
	pass
}
