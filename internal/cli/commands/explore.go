package commands

import (
	"fmt"

	"github.com/mawkler/pokedex-cli/internal"
	"github.com/mawkler/pokedex-cli/internal/cli"
)

func Explore(cfg *cli.Config, args ...string) error {
	if len(args) <= 0 {
		return fmt.Errorf("no location name passed in")
	}

	location := args[0]
	locationArea, err := internal.GetLocationArea(location)
	if err != nil {
		return fmt.Errorf("exploration failed: %s", err)
	}

	if locationArea == nil {
		return fmt.Errorf("location %s not found", location)
	}

	println("Found Pokemon:")
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
