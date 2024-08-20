package pokeapi

import (
	"fmt"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (pokemon *Pokemon) getStatsString() (stats string) {
	for _, stat := range pokemon.Stats {
		stats += fmt.Sprintf("\n  - %s: %d", stat.Stat.Name, stat.BaseStat)
	}

	return stats
}

func (pokemon *Pokemon) getTypesString() (types string) {
	for _, t := range pokemon.Types {
		types += fmt.Sprintf("\n  - %s", t.Type.Name)
	}

	return types
}

func (pokemon *Pokemon) String() string {
	stats := pokemon.getStatsString()
	types := pokemon.getTypesString()

	return fmt.Sprintf(`Name: %s
Height: %d
Weight: %d
Stats: %s
Types: %s`, pokemon.Name, pokemon.Height, pokemon.Weight, stats, types)
}
