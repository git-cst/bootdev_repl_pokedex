package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationRequest struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocation(url string) (LocationRequest, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationRequest{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationRequest{}, fmt.Errorf("status code: %d, response text: %s", res.StatusCode, res.Body)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationRequest{}, err
	}

	locations := LocationRequest{}
	err = json.Unmarshal(bodyBytes, &locations)
	if err != nil {
		return LocationRequest{}, err
	}

	return locations, nil
}
