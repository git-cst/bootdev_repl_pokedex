package pokeapi

import "encoding/json"

type LocationRequest struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocation(URL string) (LocationRequest, error) {
	bodyBytes, err := PerformGetRequest(URL)
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
