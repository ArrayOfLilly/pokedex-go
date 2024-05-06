// for parsing JSON, maps JSON data to struct
package pokeapi

type LocationAreasResp struct {
	Count int `json:"count"`              // from the response JSON
	Next *string `json:"next"`             // for example: ?limit=20&offset=20 (at start)
	Previous *string  `json:"previous"`    // for example: nil (at start), this is why a pointer to the string is better choice
	Results []struct{
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}

