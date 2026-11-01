package pokeapi

import (
	"github.com/john22n/golang-pokedex/internal/pokecache"
	"net/http"
	"time"
)

// Client
type Client struct {
	httpClient http.Client
	cache      *pokecache.Pokecache
}

// new Client
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(timeout),
	}
}
