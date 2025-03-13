package pokeapi

import "encoding/json"

type Pokemon struct {
	Id             int                   `json:"id"`
	Name           string                `json:"name"`
	Stats          map[string]StatValues `json:"stats"`
	Types          []string              `json:"type"`
	BaseExperience int                   `json:"base_experience"`
}

type StatValues struct {
	BaseValue int    `json:"base_value"`
	Effort    int    `json:"effort"`
	URL       string `json:"url"`
}

func PokemonRequest(URL string) (Pokemon, error) {
	bodyBytes, err := PerformGetRequest(URL)
	if err != nil {
		return Pokemon{}, err
	}

	// Create the raw data structure
	var rawData struct {
		BaseExperience         int    `json:"base_experience"`
		Height                 int    `json:"height"`
		ID                     int    `json:"id"`
		IsDefault              bool   `json:"is_default"`
		LocationAreaEncounters string `json:"location_area_encounters"`
		Name                   string `json:"name"`
		Order                  int    `json:"order"`
		PastAbilities          []any  `json:"past_abilities"`
		PastTypes              []any  `json:"past_types"`
		Species                struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"species"`
		Stats []struct {
			BaseStat int `json:"base_stat"`
			Effort   int `json:"effort"`
			Stat     struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"stat"`
		} `json:"stats"`
		Types []struct {
			Slot int `json:"slot"`
			Type struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"type"`
		} `json:"types"`
	}

	err = json.Unmarshal(bodyBytes, &rawData)
	if err != nil {
		return Pokemon{}, err
	}

	// Create the struct to return to the command
	pokemon := Pokemon{
		Id:             rawData.ID,
		Name:           rawData.Name,
		Stats:          make(map[string]StatValues),
		Types:          []string{},
		BaseExperience: rawData.BaseExperience,
	}

	// Process each stat adding to the stats map
	for _, statInfo := range rawData.Stats {
		statName := statInfo.Stat.Name
		pokemon.Stats[statName] = StatValues{
			BaseValue: statInfo.BaseStat,
			Effort:    statInfo.BaseStat,
			URL:       statInfo.Stat.URL,
		}
	}

	// Process each type appending to the stats slice
	for _, typeInfo := range rawData.Types {
		pokemon.Types = append(pokemon.Types, typeInfo.Type.Name)
	}

	return pokemon, nil
}
