package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
) 

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := BaseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// use cached if it exists
	if val, ok := c.cache.Get(fullURL); ok {
		locationsResp := LocationAreasResp{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationsResp, nil
	}

	// make new request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// send new HTTP request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	// don't forget to close the resp.Body, if there's one
	defer resp.Body.Close()

	// check the status code
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad statuscode: %v", resp.StatusCode)
	}

	// read data ([]Bytes) from the response body
	dat, err := io.ReadAll(resp.Body) 
	if err != nil {
		return LocationAreasResp{}, err
	}

	// parse json
	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasResp, nil
}


