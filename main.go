package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatalf("You need to provide a pokemon")
    }

    chains := GetEvolutionChain(os.Args[1])


    for _, chain := range chains {
        var str_chain strings.Builder
        for i, pokemon := range chain.Pokemons {
            if i+1 == len(chain.Pokemons){
                str_chain.WriteString(pokemon)
                break
            }
            str_chain.WriteString(fmt.Sprintf("%s->", pokemon))
        }
        fmt.Println(str_chain.String())
    }
}
