package main

import (
	"bufio"
	"fmt"
	"os"
)

func startPrompt() {
	commandMap := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		command := commandMap[input]
		command.command()
	}
}
