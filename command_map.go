package main

import (
	"errors"
	"fmt"
)

func commandMapf(c *config) error {
	locationsResp, err := c.pokeapiClient.LocationList(c.nextLocationsURL)
	if err != nil {
		return err
	}

	c.previousLocationsURL = locationsResp.Previous
	c.nextLocationsURL = locationsResp.Next

	for _, res := range locationsResp.Results {
		fmt.Println(res.Name)
	}

	return nil
}

func commandMapB(c *config) error {
	if c.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := c.pokeapiClient.LocationList(c.previousLocationsURL)
	if err != nil {
		return err
	}

	c.previousLocationsURL = locationsResp.Previous
	c.nextLocationsURL = locationsResp.Next

	for _, res := range locationsResp.Results {
		fmt.Println(res.Name)
	}
	return nil
}
