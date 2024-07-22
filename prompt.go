package main

import (
	"bufio"
	"fmt"
	"internal/pokecache"
	"os"
	"time"
)

type Config struct {
	cache  pokecache.Cache
	offset int
}

func startPrompt() {
	config := initConfig()
	go config.cache.ReapLoop()
	commandMap := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if !inputValidator(input) {
			fmt.Println("\nInvalid command, use 'help' to get valid commands!")
			continue
		}
		command := commandMap[input]
		command.command(&config)
	}
}

func inputValidator(input string) bool {
	commandMap := getCommands()
	for command := range commandMap {
		if command == input {
			return true
		}
	}
	return false
}

func initConfig() Config {
	return Config{
		cache:  pokecache.NewCache(5 * time.Minute),
		offset: 0,
	}
}
