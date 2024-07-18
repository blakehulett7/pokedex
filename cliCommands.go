package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	command     func()
}

func commandExit() {
	os.Exit(0)
}

func commandHelp() {
	fmt.Println("Welcome to the pokedex!")
	fmt.Println("")
	fmt.Println("Available Commands:")
}

func getCommands() map[string]cliCommand {
	commandMap := map[string]cliCommand{
		"exit": cliCommand{
			name:        "exit",
			description: "exit the program",
			command:     commandExit,
		},
		"help": cliCommand{
			name:        "help",
			description: "display help information",
			command:     commandHelp,
		},
	}
	return commandMap
}
