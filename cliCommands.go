package main

type cliCommand struct {
	name        string
	description string
	command     func()
}

func commandExit() {
	"exit the program!"
}

func commandHelp() {
	"help me"
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
