package main

import (
	"fmt"
	"internal/api"
	"math/rand"
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
	fmt.Println("\nExploring", argument, "...")
	url := "https://pokeapi.co/api/v2/location-area/" + argument
	entry, exists := config.cache.Map[url]
	if !exists {
		fmt.Println("Adding to cache...")
		fetchedData := api.Fetch(url)
		config.cache.Add(url, fetchedData)
		entry = config.cache.Map[url]
	}
	data := entry.Val
	if string(data) == "Not Found" {
		fmt.Println("No pokemon found... check location name!")
		fmt.Println("")
		return
	}
	fmt.Println("Found pokemon:")
	ReadEncounters(data)
	fmt.Println("")
	return
}

func commandCatch(config *Config, argument string) {
	if argument == "" {
		fmt.Println("\nAdd a pokemon to catch. Use 'help' for more info!")
		fmt.Println("")
		return
	}
	_, isCaught := config.pokedex[argument]
	if isCaught {
		fmt.Println("")
		fmt.Println(argument, "has already been caught!")
		fmt.Println("")
		return
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + argument
	entry, exists := config.cache.Map[url]
	fmt.Println("")
	if !exists {
		fmt.Println("Adding to cache...")
		fetchedData := api.Fetch(url)
		config.cache.Add(url, fetchedData)
		entry = config.cache.Map[url]
	}
	data := entry.Val
	if string(data) == "Not Found" {
		fmt.Println("Pokemon not found... check its name!")
		fmt.Println("")
		return
	}
	pokemon := ReadPokemon(data)
	fmt.Println("Throwing a pokeball at", pokemon.Name+"...")
	rngCheck := rand.Int31n(1000)
	if rngCheck < int32(pokemon.BaseExperience) {
		fmt.Println(pokemon.Name, "escaped!")
		fmt.Println("")
		return
	}
	fmt.Println(pokemon.Name, "was caught!")
	config.pokedex[pokemon.Name] = pokemon
	fmt.Println("")
	return
}

func commandInspect(config *Config, argument string) {
	if argument == "" {
		fmt.Println("\nAdd a pokemon to inspect, use 'help' for more info!")
		fmt.Println("")
		return
	}
	pokemon, isCaught := config.pokedex[argument]
	if !isCaught {
		fmt.Println("\nYou have not caught that pokemon...")
		fmt.Println("")
		return
	}
	fmt.Println("")
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Println("   -"+stat.Stat.Name+":", stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Println("   -" + pokeType.Type.Name)
	}
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
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempt to catch a pokemon",
			command:     commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Learn more about the pokemon you have caught",
			command:     commandInspect,
		},
	}
	return commandMap
}
