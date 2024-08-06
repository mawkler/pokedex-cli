package commands

import (
	"fmt"

	"github.com/mawkler/pokedex-cli/cli"
	"github.com/mawkler/pokedex-cli/internal"
)

func updateConfig(page *internal.LocationAreasPage, cfg *cli.Config) {
	if page.Next != nil {
		cfg.Next = page.Next
	}

	if page.Previous != nil {
		cfg.Previous = page.Previous
	}
}

func mapAndUpdateConfig(url *string, cfg *cli.Config) error {
	if url == nil {
		defaultURL := "https://pokeapi.co/api/v2/location-area"
		url = &defaultURL
	}

	page, err := internal.GetLocationAreas(*url)
	if err != nil {
		return err
	}

	updateConfig(page, cfg)

	page.Print()
	return nil
}

func Map(cfg *cli.Config) error {
	return mapAndUpdateConfig(cfg.Next, cfg)
}

func Mapb(cfg *cli.Config) error {
	if cfg.Previous == nil {
		fmt.Println("Already on the first page")
		return nil
	}

	return mapAndUpdateConfig(cfg.Previous, cfg)
}
