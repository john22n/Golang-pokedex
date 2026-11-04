package main

import (
	"errors"
	"fmt"
)

func commandPokedex(c *config, args ...string) error {
	if len(c.pokedex) == 0 {
		return errors.New("empy pokedex")
	}
	for _, v := range c.pokedex {
		fmt.Printf(" - %s\n", v.Name)

	}
	return nil
}
