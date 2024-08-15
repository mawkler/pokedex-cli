package pokeapi

import "fmt"

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func (client *Client) GetPokmeon(name string) (*Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", name)
	data, err := client.get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get pokemon %s: %s", name, err)
	}

	if data == nil {
		return nil, nil
	}

	pokemon, err := unmarshal[Pokemon](*data)
	if err != nil {
		return nil, fmt.Errorf("failed to get pokemon %s: %s", name, err)
	}

	return pokemon, nil
}
