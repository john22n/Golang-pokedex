package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationList(pageURL *string) (RespShallowLocations, error) {
	locationResp := RespShallowLocations{}

	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	entry, ok := c.cache.Get(url)
	if ok {
		json.Unmarshal(entry, &locationResp)
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationResp, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationResp, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationResp, err
	}

	c.cache.Add(url, dat)

	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return locationResp, err
	}

	return locationResp, nil
}
