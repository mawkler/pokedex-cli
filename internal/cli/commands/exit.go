package commands

import (
	"os"

	"github.com/mawkler/pokedex-cli/internal/cli"
)

func Exit(_ *cli.Config, _ ...string) error {
	println("Bye!")
	os.Exit(0)
	return nil
}
