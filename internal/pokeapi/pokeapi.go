package pokeapi

import (
	"net/http"
	"time"

	"github.com/ArrayOfLilly/pokedexcli/internal/pokecache"
)

// common part of URL Path, without any endpoint
const BaseURL = "https://pokeapi.co/api/v2"

// Own version, to keep a timer
type Client struct{
	cache      pokecache.Cache
	httpClient http.Client
} 

// Constructor-like function, for export (Caps...)
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client {
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			// after 1 min, http request is killed
			Timeout: time.Minute,
		},
	}
}

