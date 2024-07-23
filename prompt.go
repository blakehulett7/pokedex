package main

import (
	"bufio"
	"fmt"
	"internal/pokecache"
	"os"
	"strings"
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
	for {
		input, argument := ReadInput()
		if !inputValidator(input) {
			fmt.Println("\nInvalid command, use 'help' to get valid commands!")
			continue
		}
		command := commandMap[input]
		command.command(&config, argument)
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

func ReadInput() (string, string) {
	prompt := bufio.NewScanner(os.Stdin)
	fmt.Print("pokedex > ")
	prompt.Scan()
	input := prompt.Text()
	fmt.Println("input:", input)
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	command := scanner.Text()
	scanner.Scan()
	argument := scanner.Text()
	return command, argument
}
