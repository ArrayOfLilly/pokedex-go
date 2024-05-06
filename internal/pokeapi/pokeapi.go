package pokeapi

import (
	"net/http"
	"time"
)

// common part of URL Path, without any endpoint
const BaseURL = "https://pokeapi.co/api/v2"

// Own version, to keep a timer
type Client struct{
	httpClient http.Client
} 

// Constructor-like function, for export (Caps...)
func NewClient() Client {
	return Client {
		httpClient: http.Client{
			// after 1 min, http request is killed
			Timeout: time.Minute,
		},
	}
}

