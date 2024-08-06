package commands

import (
	"github.com/mawkler/pokedex-cli/cli"
)

type Command struct {
	callback    func(*cli.Config) error
	name        string
	description string
}

func (cmd *Command) Run(cfg *cli.Config) error {
	return cmd.callback(cfg)
}

func NewCLICommandMap() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    Help,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    Exit,
		},
		"map": {
			name:        "map",
			description: "List 20 next location areas",
			callback:    Map,
		},
		"mapb": {
			name:        "mapb",
			description: "List 20 previous location areas",
			callback:    Mapb,
		},
	}
}
