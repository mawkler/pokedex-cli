package commands

import (
	"os"

	"github.com/mawkler/pokedex-cli/cli"
)

func Exit(_ *cli.Config) error {
	println("Bye!")
	os.Exit(0)
	return nil
}
