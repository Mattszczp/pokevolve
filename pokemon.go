package main

import (
	"log"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

type EvolutionChain struct {
	Pokemons []string
}

//TODO: Replace panicking with returning HTTP errors
func GetEvolutionChain(pokemon_name string) []EvolutionChain {
	pokemon, err := pokeapi.PokemonSpecies(pokemon_name)
	if err != nil {
		log.Fatalf("Couldn't find a pokemon named: %s", pokemon_name)
	}

	chain := pokemon.EvolutionChain.URL
	chain_id := strings.Split(chain, "/")[6]

	evo_chain, err := pokeapi.EvolutionChain(chain_id)

	if err != nil {
		log.Fatalf("Couldn't find th evolution chain for %s", pokemon_name)
	}

	var chains []EvolutionChain

	for _, evolution := range evo_chain.Chain.EvolvesTo {

		var chain EvolutionChain

		chain.Pokemons = append(chain.Pokemons, evo_chain.Chain.Species.Name, evolution.Species.Name)

		if len(evolution.EvolvesTo) > 0 {
			chain.Pokemons = append(chain.Pokemons, evolution.EvolvesTo[0].Species.Name)
		}

		chains = append(chains, chain)
	}

	return chains
}
