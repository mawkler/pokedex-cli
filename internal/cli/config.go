package cli

import "github.com/mawkler/pokedex-cli/internal/pokeapi"

type Config struct {
	Next     *string
	Previous *string
	Client   pokeapi.Client
}

func NewConfig(client pokeapi.Client) Config {
	return Config{Client: client}
}

func (cfg *Config) setNext(next *string) {
	if next == nil {
		return
	}

	cfg.Next = next
}
