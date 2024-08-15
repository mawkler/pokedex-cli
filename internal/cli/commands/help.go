package commands

import (
	"fmt"

	"github.com/mawkler/pokedex-cli/internal/cli"
)

func Help(config *cli.Config, commands ...string) error {
	commandMap := NewCLICommandMap()

	if len(commands) == 1 {
		command := commands[0]

		if cmd, ok := commandMap[command]; !ok {
			fmt.Printf("command %s not found\n", cmd)
		} else {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		}

		return nil
	}

	for _, command := range commandMap {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
