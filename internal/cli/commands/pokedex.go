package commands

import "github.com/mawkler/pokedex-cli/internal/cli"

func Pokedex(cfg *cli.Config, args ...string) error {
	pokemons := cfg.Pokedex.GetAllNames()

	if len(pokemons) == 0 {
		println("Your Pokedex is currently empty")
		return nil
	}

	println("Your Pokedex:")
	for _, pokemon := range pokemons {
		println("  -", pokemon)
	}

	return nil
}
