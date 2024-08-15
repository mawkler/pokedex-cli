package pokeapi

import (
	"fmt"
)

type LocationAreasPage struct {
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	Count int `json:"count"`
}

type LocationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetLocationAreas(url string) (*LocationAreasPage, error) {
	page, err := get[LocationAreasPage](url)
	if err != nil {
		return nil, fmt.Errorf("failed to get location areas: %s", err)
	}

	return page, nil
}

func (page *LocationAreasPage) Print() {
	for _, area := range page.Results {
		fmt.Println(area.Name)
	}
}

func GetLocationArea(locationArea string) (*LocationArea, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", locationArea)
	location, err := get[LocationArea](url)
	if err != nil {
		return nil, fmt.Errorf("failed to get location area %s: %s", locationArea, err)
	}

	return location, nil
}
