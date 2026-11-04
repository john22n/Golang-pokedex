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

	pokemon := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	caught, err := c.pokeapiClient.CaughtPokemon(pokemon)

	if err != nil {
		return err
	}

	r := rand.IntN(int(caught.BaseExperience))

	if r > 40 {
		fmt.Printf("%s was caught!\n", caught.Name)
	} else {
		fmt.Printf("%s escaped\n", pokemon)
		return nil
	}

	c.pokedex[caught.Name] = caught
	return nil
}
