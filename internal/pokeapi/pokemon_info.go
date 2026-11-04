package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonInfo(pokemon string) (Pokemon, error) {
	pokemonInfo := Pokemon{}

	url := baseURL + "/pokemon/" + pokemon

	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &pokemonInfo)
		if err != nil {
			return pokemonInfo, err
		}
		return pokemonInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemonInfo, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokemonInfo, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemonInfo, err
	}

	err = json.Unmarshal(data, &pokemonInfo)

	if err != nil {
		return pokemonInfo, err
	}

	c.cache.Add(url, data)

	return pokemonInfo, nil
}
