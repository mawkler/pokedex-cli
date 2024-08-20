package pokedex

import (
	types "github.com/mawkler/pokedex-cli/internal/pokeapi/types"
)

type Pokedex struct {
	pokemon map[string]types.Pokemon
}

func NewPokedex() Pokedex {
	return Pokedex{pokemon: map[string]types.Pokemon{}}
}

func (pdx *Pokedex) Add(name string, pokemon types.Pokemon) {
	pdx.pokemon[name] = pokemon
}

func (pdx *Pokedex) Get(name string) *types.Pokemon {
	if pokemon, exists := pdx.pokemon[name]; exists {
		return &pokemon
	} else {
		return nil
	}
}
