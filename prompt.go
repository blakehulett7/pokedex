package main

import (
	"bufio"
	"fmt"
	"os"
)

type Config struct {
	offset int
}

func startPrompt() {
	config := Config{offset: 0}
	commandMap := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if !inputValidator(input) {
			fmt.Println("\nInvalid command, use 'help' to get valid commands!\n")
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
