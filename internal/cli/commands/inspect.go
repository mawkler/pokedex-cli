package commands

import (
	"fmt"

	"github.com/mawkler/pokedex-cli/internal/cli"
)

func Inspect(cfg *cli.Config, args ...string) error {
	pokemonName := args[0]
	pokemon := cfg.Pokedex.Get(pokemonName)

	if pokemon == nil {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("%s\n", pokemon)

	return nil
}
