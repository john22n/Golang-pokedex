package main

import (
	"github.com/john22n/golang-pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	cfg := config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
	}

	startRepl(&cfg)
}
