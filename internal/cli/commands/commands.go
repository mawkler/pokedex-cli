package commands

import (
	"github.com/mawkler/pokedex-cli/internal/cli"
)

type Command struct {
	callback    func(*cli.Config, ...string) error
	name        string
	description string
}

func (cmd *Command) Run(cfg *cli.Config, input ...string) error {
	return cmd.callback(cfg, input...)
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
		"explore": {
			name:        "explore",
			description: "Explore a given location",
			callback:    Explore,
		},
	}
}
