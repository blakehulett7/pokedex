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
	Map   map[string]cacheEntry
	Mutex *sync.Mutex
}

func NewCache(interval time.Duration) {
	cache := Cache{
		Map:   map[string]cacheEntry{},
		Mutex: &sync.Mutex{},
	}
	cache.reapLoop(interval)
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

func (c Cache) reapLoop(intervalSeconds time.Duration) {
	interval := intervalSeconds * 1000000000
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		fmt.Println("Ticks Complete")
	}
}
