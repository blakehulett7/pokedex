package main

import (
	"fmt"
	"internal/api"
	"os"
)

type cliCommand struct {
	name        string
	description string
	command     func(Config)
}

func commandExit(Config) {
	os.Exit(0)
}

func commandHelp(Config) {
	fmt.Println("\nWelcome to the pokedex!")
	fmt.Println("")
	fmt.Println("-- Available Commands --")
	fmt.Println("")
	commandMap := getCommands()
	for _, command := range commandMap {
		fmt.Println(command.name + ": " + command.description)
	}
	fmt.Println("")
}

func commandMap(Config) {
	api.Fetch("https://pokeapi.co/api/v2/location-area" + "?offset=0")
	return
}

func commandMapBack(Config) {
	return
}

func getCommands() map[string]cliCommand {
	commandMap := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the program",
			command:     commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			command:     commandHelp,
		},
		"map": {
			name:        "map",
			description: "Explore the world of pokemon",
			command:     commandMap,
		},
		"mapback": {
			name:        "mapback",
			description: "Go to previous page of the map",
			command:     commandMapBack,
		},
	}
	return commandMap
}
