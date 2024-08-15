package pokedex

import (
	"github.com/mawkler/pokedex-cli/internal/pokeapi"
)

type Pokedex struct {
	pokemon map[string]pokeapi.Pokemon
}

func NewPokedex() Pokedex {
	return Pokedex{pokemon: map[string]pokeapi.Pokemon{}}
}

func (pdx *Pokedex) Add(name string, pokemon pokeapi.Pokemon) {
	pdx.pokemon[name] = pokemon
}

func (pdx *Pokedex) Get(name string) *pokeapi.Pokemon {
	if pokemon, exists := pdx.pokemon[name]; exists {
		return &pokemon
	} else {
		return nil
	}
}
