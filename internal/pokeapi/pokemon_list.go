package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonList(name string) (Pokemons, error) {
	pokemonsResp := Pokemons{}

	url := baseURL + "/location-area/" + name

	entry, ok := c.cache.Get(url)
	if ok {
		json.Unmarshal(entry, &pokemonsResp)
		return pokemonsResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemons{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemons{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemons{}, nil
	}
	c.cache.Add(url, data)

	err = json.Unmarshal(data, &pokemonsResp)
	if err != nil {
		return Pokemons{}, nil
	}

	return pokemonsResp, nil
}
