package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
) 

func (c *Client) DescribePokemon(pokemonName string) (PokemonResp, error) {
	endpoint := "/pokemon" + "/" + pokemonName
	fullURL := BaseURL + endpoint
	if len(pokemonName) == 0 {
		return PokemonResp{}, fmt.Errorf("missing pokemon name")
	}

	// use cached if it exists
	if val, ok := c.cache.Get(fullURL); ok {
		pokemonResp := PokemonResp{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return PokemonResp{}, err
		}

		return pokemonResp, nil
	}

	// make new request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResp{}, err
	}

	// send new HTTP request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResp{}, err
	}
	// don't forget to close the resp.Body, if there's one
	defer resp.Body.Close()

	// check the status code
	if resp.StatusCode > 399 {
		return PokemonResp{}, fmt.Errorf("bad statuscode: %v", resp.StatusCode)
	}

	// read data ([]Bytes) from the response body
	dat, err := io.ReadAll(resp.Body) 
	if err != nil {
		return PokemonResp{}, err
	}

	// parse json
	pokemonResp := PokemonResp{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonResp{}, err
	}

	return pokemonResp, nil
}