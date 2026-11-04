package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonListLocation(name string) (Location, error) {
	pokemonsResp := Location{}

	url := baseURL + "/location-area/" + name

	entry, ok := c.cache.Get(url)
	if ok {
		json.Unmarshal(entry, &pokemonsResp)
		return pokemonsResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, nil
	}
	c.cache.Add(url, data)

	err = json.Unmarshal(data, &pokemonsResp)
	if err != nil {
		return Location{}, nil
	}

	return pokemonsResp, nil
}
