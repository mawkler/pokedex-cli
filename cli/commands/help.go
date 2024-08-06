package commands

import (
	"fmt"

	"github.com/mawkler/pokedex-cli/cli"
)

func Help(config *cli.Config) error {
	for _, command := range NewCLICommandMap() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
