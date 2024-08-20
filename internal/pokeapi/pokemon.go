package pokeapi

import (
	"fmt"

	types "github.com/mawkler/pokedex-cli/internal/pokeapi/types"
)

func (client *Client) GetPokmeon(name string) (*types.Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s/", client.baseUrl, name)
	data, err := client.get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get pokemon %s: %s", name, err)
	}

	if data == nil {
		return nil, nil
	}

	pokemon, err := unmarshal[types.Pokemon](*data)
	if err != nil {
		return nil, fmt.Errorf("failed to get pokemon %s: could not unmarshal: %s", name, err)
	}

	return pokemon, nil
}
