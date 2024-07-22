package main

import (
	"fmt"
	"internal/pokecache"
)

func main() {
	fmt.Println("Christ is King!")
	cache := pokecache.NewCache(20)
	cache.ReapLoop()
}
