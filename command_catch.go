package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon to catch")
	}

	min := 0
	max := 608
	r := rand.IntN(max-min+1) + min

	pokemon := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	caught, err := c.pokeapiClient.CaughtPokemon(pokemon)

	if err != nil {
		return err
	}

	if r > int(caught.BaseExperience) {
		fmt.Printf("%s was caught!\n", caught.Name)
		c.pokedex[pokemon] = Pokemon{
			name:           pokemon,
			baseExperience: caught.BaseExperience,
		}
	} else {
		fmt.Printf("%s escaped\n", pokemon)
	}

	return nil
}
