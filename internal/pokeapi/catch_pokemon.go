package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CaughtPokemon(name string) (CaughtPokemon, error) {
	caughtResponse := CaughtPokemon{}

	url := baseURL + "/pokemon/" + name

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return caughtResponse, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return caughtResponse, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return caughtResponse, err
	}

	err = json.Unmarshal(data, &caughtResponse)
	if err != nil {
		return caughtResponse, err
	}

	return caughtResponse, nil
}
