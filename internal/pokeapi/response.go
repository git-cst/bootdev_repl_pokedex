package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func PerformGetRequest(URL string) ([]byte, error) {
	res, err := http.Get(URL)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("status code: %d\n, response text: %s", res.StatusCode, res.Body)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return bodyBytes, nil
}
