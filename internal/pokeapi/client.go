package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func get[T any](url string) (*T, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get failed: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("get failed: could not read body: %s", err)
	}

	if res.StatusCode == 404 {
		return nil, nil
	}

	if res.StatusCode >= 300 {
		msg := "get failed: response failed with status code: %d and body: %s"
		return nil, fmt.Errorf(msg, res.StatusCode, body)
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("get failed: failed to unmarshal body: %s", err)
	}

	return &result, nil
}
