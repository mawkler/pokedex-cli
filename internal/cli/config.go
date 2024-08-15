package cli

import (
	"github.com/mawkler/pokedex-cli/internal/pokeapi"
	"github.com/mawkler/pokedex-cli/internal/pokedex"
)

type Config struct {
	Next     *string
	Previous *string
	Client   pokeapi.Client
	Pokedex  pokedex.Pokedex
}

func NewConfig(client pokeapi.Client, pokedex pokedex.Pokedex) Config {
	return Config{Client: client, Pokedex: pokedex}
}

func (cfg *Config) setNext(next *string) {
	if next == nil {
		return
	}

	cfg.Next = next
}
