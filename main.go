package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatalf("You need to provide a pokemon")
    }
    pokemon_name := os.Args[1]
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


    for _, evolution := range evo_chain.Chain.EvolvesTo {

        if len(evolution.EvolvesTo) > 0 {
            fmt.Println(evo_chain.Chain.Species.Name,"->",evolution.Species.Name, "->", evolution.EvolvesTo[0].Species.Name)
            break
        }
        fmt.Println(evo_chain.Chain.Species.Name,"->",evolution.Species.Name)
    }


}
