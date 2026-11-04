package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon")
	}

	pokemon := args[0]

	info, ok := c.pokedex[pokemon]
	if !ok {
		return errors.New("you have not caught this pokemon")
	}
	fmt.Println("Name:", info.Name)
	fmt.Println("Height:", info.Height)
	fmt.Println("Weight:", info.Weight)
	fmt.Println("Stats:")
	for _, stat := range info.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range info.Types {
		fmt.Println(" -", typeInfo.Type.Name)
	}
	return nil
}
