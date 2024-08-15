package commands

import (
	"errors"
	"fmt"

	"github.com/mawkler/pokedex-cli/internal/cli"
	"github.com/mawkler/pokedex-cli/internal/pokeapi"
)

func updateConfig(page *pokeapi.LocationAreasPage, cfg *cli.Config) {
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

	page, err := cfg.Client.GetLocationAreas(*url)
	if err != nil {
		return err
	}

	if page == nil {
		return errors.New("page not found")
	}

	updateConfig(page, cfg)

	page.Print()
	return nil
}

func Map(cfg *cli.Config, _ ...string) error {
	if cfg.Next == nil && cfg.Previous != nil {
		return fmt.Errorf("already on the last page")
	}
	if cfg.Next != nil {
	}
	err := mapAndUpdateConfig(cfg.Next, cfg)
	return err
}

func Mapb(cfg *cli.Config, _ ...string) error {
	if cfg.Previous == nil {
		return fmt.Errorf("already on the first page")
	}

	return mapAndUpdateConfig(cfg.Previous, cfg)
}
