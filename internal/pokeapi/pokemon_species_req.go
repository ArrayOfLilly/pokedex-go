package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
) 

func (c *Client) DescribePokemonSpecies(pokemonID string) (PokemonSpeciesResp, error) {
	endpoint := "/pokemon-species" + "/" + pokemonID
	fullURL := BaseURL + endpoint
	if len(pokemonID) == 0 {
		return PokemonSpeciesResp{}, fmt.Errorf("missing pokemon ID")
	}

	// use cached if it exists
	if val, ok := c.cache.Get(fullURL); ok {
		pokemonSpeciesResp := PokemonSpeciesResp{}
		err := json.Unmarshal(val, &pokemonSpeciesResp)
		if err != nil {
			return PokemonSpeciesResp{}, err
		}

		return pokemonSpeciesResp, nil
	}

	// make new request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonSpeciesResp{}, err
	}

	// send new HTTP request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonSpeciesResp{}, err
	}
	// don't forget to close the resp.Body, if there's one
	defer resp.Body.Close()

	// check the status code
	if resp.StatusCode > 399 {
		return PokemonSpeciesResp{}, fmt.Errorf("bad statuscode: %v", resp.StatusCode)
	}

	// read data ([]Bytes) from the response body
	dat, err := io.ReadAll(resp.Body) 
	if err != nil {
		return PokemonSpeciesResp{}, err
	}

	// parse json
	pokemonSpeciesResp := PokemonSpeciesResp{}
	err = json.Unmarshal(dat, &pokemonSpeciesResp)
	if err != nil {
		return PokemonSpeciesResp{}, err
	}

	return pokemonSpeciesResp, nil
}