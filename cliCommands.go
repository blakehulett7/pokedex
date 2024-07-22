package main

import (
	"fmt"
	"internal/api"
	"os"
)

type cliCommand struct {
	name        string
	description string
	command     func(*Config)
}

func commandExit(config *Config) {
	os.Exit(0)
}

func commandHelp(config *Config) {
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

func commandMap(config *Config) {
	param := fmt.Sprintf("?offset=%v", config.offset)
	url := "https://pokeapi.co/api/v2/location-area" + param
	api.Fetch(url)
	config.offset += 20
	return
}

func commandMapBack(config *Config) {
	config.offset -= 40
	param := fmt.Sprintf("?offset=%v", config.offset)
	api.Fetch("https://pokeapi.co/api/v2/location-area" + param)
	config.offset += 20
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
