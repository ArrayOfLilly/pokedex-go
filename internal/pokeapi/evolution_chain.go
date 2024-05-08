package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
) 

func (c *Client) GetEvolutionChain(url string) (EvolutionChainResp, error) {
	fullURL := url
	if len(url) == 0 {
		return EvolutionChainResp{}, fmt.Errorf("invalid url")
	}

	// use cached if it exists
	if val, ok := c.cache.Get(fullURL); ok {
		evolutionChainResp := EvolutionChainResp{}
		err := json.Unmarshal(val, &evolutionChainResp)
		if err != nil {
			return EvolutionChainResp{}, err
		}

		return evolutionChainResp, nil
	}

	// make new request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return EvolutionChainResp{}, err
	}

	// send new HTTP request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return EvolutionChainResp{}, err
	}
	// don't forget to close the resp.Body, if there's one
	defer resp.Body.Close()

	// check the status code
	if resp.StatusCode > 399 {
		return EvolutionChainResp{}, fmt.Errorf("bad statuscode: %v", resp.StatusCode)
	}

	// read data ([]Bytes) from the response body
	dat, err := io.ReadAll(resp.Body) 
	if err != nil {
		return EvolutionChainResp{}, err
	}

	// parse json
	evolutionChainResp := EvolutionChainResp{}
	err = json.Unmarshal(dat, &evolutionChainResp)
	if err != nil {
		return EvolutionChainResp{}, err
	}

	return evolutionChainResp, nil
}