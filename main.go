package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mtslzr/pokeapi-go"
)

func main() {
    pokemon_name := os.Args[1]
    pokemon, err := pokeapi.Pokemon(pokemon_name)
    if err != nil {
        log.Fatalf("Couldn't find a pokemon named: %s", pokemon_name)
    }

    pokemon_id := strconv.Itoa(pokemon.ID)

    fmt.Println("The id of", pokemon_name , "is:", pokemon_id)

    if err != nil {
        log.Fatalf("Couldn't convert pokemon id to int")
    }
}
