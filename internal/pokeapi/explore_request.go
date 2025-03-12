package pokeapi

import "encoding/json"

type PokemonAttributes struct {
	MinLevel int `json:"min_level"`
	MaxLevel int `json:"max_level"`
	Chance   int `json:"chance"`
}

type Pokemon struct {
	Name              string                       `json:"name"`
	URL               string                       `json:"url"`
	VersionAttributes map[string]PokemonAttributes `json:"attributes"`
}

type ExploreRequest struct {
	LocationName string             `json:"name"`
	LocationURL  string             `json:"url"`
	Pokemon      map[string]Pokemon `json:"pokemon"`
}

func ExploreLocation(url string) (ExploreRequest, error) {
	bodyBytes, err := PerformGetRequest(url)
	if err != nil {
		return ExploreRequest{}, err
	}

	// Create the raw data structure
	var rawData struct {
		Location struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"location"`
		PokemonEncounters []struct {
			Pokemon struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"pokemon"`
			VersionDetails []struct {
				EncounterDetails []struct {
					Chance   int `json:"chance"`
					MaxLevel int `json:"max_level"`
					MinLevel int `json:"min_level"`
					Method   struct {
						Name string `json:"name"`
					} `json:"method"`
				} `json:"encounter_details"`
				Version struct {
					Name string `json:"name"`
				} `json:"version"`
			} `json:"version_details"`
		} `json:"pokemon_encounters"`
	}

	err = json.Unmarshal(bodyBytes, &rawData)
	if err != nil {
		return ExploreRequest{}, err
	}

	// Create the struct to return to the command
	result := ExploreRequest{
		LocationName: rawData.Location.Name,
		LocationURL:  rawData.Location.URL,
		Pokemon:      make(map[string]Pokemon),
	}

	// Process each encounter creating the pokemon
	for _, encounter := range rawData.PokemonEncounters {
		pokemon := Pokemon{
			Name:              encounter.Pokemon.Name,
			URL:               encounter.Pokemon.URL,
			VersionAttributes: make(map[string]PokemonAttributes),
		}

		// Process each version's details
		for _, versionDetail := range encounter.VersionDetails {
			versionName := versionDetail.Version.Name

			if len(versionDetail.EncounterDetails) > 0 {
				detail := versionDetail.EncounterDetails[0]
				pokemon.VersionAttributes[versionName] = PokemonAttributes{
					MinLevel: detail.MinLevel,
					MaxLevel: detail.MaxLevel,
					Chance:   detail.Chance,
				}
			}
		}

		result.Pokemon[pokemon.Name] = pokemon
	}

	locationInfo := ExploreRequest{}
	err = json.Unmarshal(bodyBytes, &locationInfo)
	if err != nil {
		return ExploreRequest{}, err
	}

	return locationInfo, nil
}
