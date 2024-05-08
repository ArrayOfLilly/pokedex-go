package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
) 

func (c *Client) ListPokemonsPerAreas(areaName string) (PokemonPerAreaResp, error) {
	endpoint := "/location-area" + "/" + areaName
	fullURL := BaseURL + endpoint
	if len(areaName) == 0 {
		return PokemonPerAreaResp{}, fmt.Errorf("missing area name")
	}

	// use cached if it exists
	if val, ok := c.cache.Get(fullURL); ok {
		pokemonPerAreaResp := PokemonPerAreaResp{}
		err := json.Unmarshal(val, &pokemonPerAreaResp)
		if err != nil {
			return PokemonPerAreaResp{}, err
		}

		return pokemonPerAreaResp, nil
	}

	// make new request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonPerAreaResp{}, err
	}

	// send new HTTP request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonPerAreaResp{}, err
	}
	// don't forget to close the resp.Body, if there's one
	defer resp.Body.Close()

	// check the status code
	if resp.StatusCode > 399 {
		return PokemonPerAreaResp{}, fmt.Errorf("bad statuscode: %v", resp.StatusCode)
	}

	// read data ([]Bytes) from the response body
	dat, err := io.ReadAll(resp.Body) 
	if err != nil {
		return PokemonPerAreaResp{}, err
	}

	// parse json
	pokemonPerAreaResp := PokemonPerAreaResp{}
	err = json.Unmarshal(dat, &pokemonPerAreaResp)
	if err != nil {
		return PokemonPerAreaResp{}, err
	}

	return pokemonPerAreaResp, nil
}