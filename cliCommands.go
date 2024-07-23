package main

import (
	"fmt"
	"internal/api"
	"os"
)

type cliCommand struct {
	name        string
	description string
	command     func(*Config, string)
}

func commandExit(config *Config, argument string) {
	os.Exit(0)
}

func commandHelp(config *Config, argument string) {
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

func commandMap(config *Config, argument string) {
	param := fmt.Sprintf("?offset=%v", config.offset)
	url := "https://pokeapi.co/api/v2/location-area" + param
	entry, exists := config.cache.Map[url]
	if !exists {
		fmt.Println("Adding to cache...")
		fetchedData := api.Fetch(url)
		config.cache.Add(url, fetchedData)
		entry = config.cache.Map[url]
	}
	data := entry.Val
	ReadLocations(data)
	config.offset += 20
	return
}

func commandMapBack(config *Config, argument string) {
	config.offset -= 40
	param := fmt.Sprintf("?offset=%v", config.offset)
	url := "https://pokeapi.co/api/v2/location-area" + param
	entry, exists := config.cache.Map[url]
	if !exists {
		fmt.Println("Adding to cache...")
		fetchedData := api.Fetch(url)
		config.cache.Add(url, fetchedData)
		entry = config.cache.Map[url]
	}
	data := entry.Val
	ReadLocations(data)
	config.offset += 20
	return
}

func commandExplore(config *Config, argument string) {
	if argument == "" {
		fmt.Println("\nAdd a location to explore. Use 'help' for more info!")
		fmt.Println("")
		return
	}
	fmt.Println("Exploring", argument)
	fmt.Println("Found pokemon:")
	url := "https://pokeapi.co/api/v2/location-area/" + argument
	data := api.Fetch(url)
	ReadEncounters(data)
	fmt.Println("")
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
			description: "Find locations in the wonderful world of pokemon",
			command:     commandMap,
		},
		"mapback": {
			name:        "mapback",
			description: "Go to previous page of the map",
			command:     commandMapBack,
		},
		"explore": {
			name:        "explore <location>",
			description: "Explore the various locations in the wonderful world of pokemon",
			command:     commandExplore,
		},
	}
	return commandMap
}
