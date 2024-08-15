package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func get[T any](url string) (*T, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get failed: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("get failed: could not read body: %s", err)
	}

	if res.StatusCode == 404 {
		return nil, nil
	}

	if res.StatusCode >= 300 {
		msg := "get failed: response failed with status code: %d and body: %s"
		return nil, fmt.Errorf(msg, res.StatusCode, body)
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("get failed: failed to unmarshal body: %s", err)
	}

	return &result, nil
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
