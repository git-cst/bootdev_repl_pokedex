package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func PerformGetRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("status code: %d, response text: %s", res.StatusCode, res.Body)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return bodyBytes, nil
}
