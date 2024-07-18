package main

import (
	"fmt"
	"internal/api"
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

func commandMap() {
	api.Fetch("https://pokeapi.co/api/v2/location-area")
	return
}

func commandMapBack() {
	return
}

func getCommands() map[string]cliCommand {
	commandMap := map[string]cliCommand{
		"exit": cliCommand{
			name:        "exit",
			description: "Exit the program",
			command:     commandExit,
		},
		"help": cliCommand{
			name:        "help",
			description: "Display a help message",
			command:     commandHelp,
		},
		"map": cliCommand{
			name:        "map",
			description: "Explore the world of pokemon",
			command:     commandMap,
		},
		"mapback": cliCommand{
			name:        "mapback",
			description: "Go to previous page of the map",
			command:     commandMapBack,
		},
	}
	return commandMap
}
