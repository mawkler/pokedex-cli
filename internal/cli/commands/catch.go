package commands

import (
	"fmt"
	"math/rand"

	"github.com/mawkler/pokedex-cli/internal/cli"
)

func Catch(cfg *cli.Config, args ...string) error {
	if len(args) <= 0 {
		return fmt.Errorf("no pokemon pokemon passed in")
	}

	pokemonName := args[0]
	pokemon, err := cfg.Client.GetPokmeon(pokemonName)
	if err != nil {
		return fmt.Errorf("failed to catch %s: %s", pokemonName, err)
	}

	if pokemon == nil {
		return fmt.Errorf("pokemon %s not found", pokemonName)
	}

	if pokemon.BaseExperience > rand.Intn(500) {
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
