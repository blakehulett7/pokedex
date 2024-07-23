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
	scanner.Split(bufio.ScanWords)
	for {
		input, argument := ReadInput(scanner)
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

func ReadInput(scanner *bufio.Scanner) (string, string) {
	fmt.Print("pokedex > ")
	scanner.Scan()
	input := scanner.Text()
	scanner.Scan()
	argument := scanner.Text()
	fmt.Println(input, "arg:", argument)
	argument = ""
	return input, argument
}
