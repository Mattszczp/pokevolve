package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/mtslzr/pokeapi-go"
)

type Pokemon struct {
	Id        int
	Name      string
	SplashUrl string
}

type EvolutionChain struct {
	Pokemons []Pokemon
}

func GetEvolutionChain(pokemon_name string) ([]EvolutionChain, error) {

	if pokemon_name == "" {
		return nil, fmt.Errorf("you have to provide a pokemon name")
	}

	pokemon, err := pokeapi.PokemonSpecies(pokemon_name)
	if err != nil {
		return nil, fmt.Errorf("couldn't find a pokemon named: %s", pokemon_name)
	}

	chain := pokemon.EvolutionChain.URL
	chain_id := strings.Split(chain, "/")[6]

	evo_chain, err := pokeapi.EvolutionChain(chain_id)

	if err != nil {
		return nil, fmt.Errorf("couldn't find th evolution chain for %s", pokemon_name)
	}

	var chains []EvolutionChain

	for _, evolution := range evo_chain.Chain.EvolvesTo {

		var chain EvolutionChain

		parent_pokemon, err := pokeapi.Pokemon(evo_chain.Chain.Species.Name)

		if err != nil {
			return nil, fmt.Errorf("couldn't find parent pokemon")
		}

		evo_pokemon, err := pokeapi.Pokemon(evolution.Species.Name)
		if err != nil {
			return nil, fmt.Errorf("couldn't get evolution")
		}

		caser := cases.Title(language.AmericanEnglish)

		chain.Pokemons = append(chain.Pokemons, Pokemon{
			Id:        parent_pokemon.ID,
			Name:      caser.String(parent_pokemon.Name),
			SplashUrl: fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/%d.png", parent_pokemon.ID),
		}, Pokemon{
			Id:        evo_pokemon.ID,
			Name:      caser.String(evo_pokemon.Name),
			SplashUrl: fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/%d.png", evo_pokemon.ID),
		})

		if len(evolution.EvolvesTo) > 0 {
			last_evo_pokemon, err := pokeapi.Pokemon(evolution.EvolvesTo[0].Species.Name)
			if err != nil {
				return nil, fmt.Errorf("couldn't get evolution")
			}
			chain.Pokemons = append(chain.Pokemons, Pokemon{
				Id:        last_evo_pokemon.ID,
				Name:      caser.String(last_evo_pokemon.Name),
				SplashUrl: fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/%d.png", last_evo_pokemon.ID),
			})
		}

		chains = append(chains, chain)
	}

	return chains, nil
}
