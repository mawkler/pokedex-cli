package cli

import (
	"fmt"
	"os"
)

type Command struct {
	callback    func() error
	name        string
	description string
}

func (cmd *Command) Run() error {
	return cmd.callback()
}

func commandHelp() error {
	for _, command := range NewCLICommandMap() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandExit() error {
	println("Bye!")
	os.Exit(0)
	return nil
}

func NewCLICommandMap() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
