package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Christ is King!")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("pokedex > ")
	scanner.Scan()
	command := scanner.Text()
	commandMap := getCommands()
	fmt.Println(commandMap[command])
}
