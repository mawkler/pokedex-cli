package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mawkler/pokedex-cli/internal/cache"
)

type Client struct {
	baseUrl string
	client  http.Client
	cache   cache.Cache
}

func NewClient(baseUrl string, client http.Client, cache cache.Cache) Client {
	return Client{baseUrl, client, cache}
}

func unmarshal[T any](data []byte) (*T, error) {
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("get failed: failed to unmarshal body: %s", err)
	}

	return &result, nil
}

func (client *Client) get(url string) (*[]byte, error) {
	entry, exists := client.cache.Get(url)
	if exists {
		return &entry, nil
	}

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

	client.cache.Add(url, body)

	return &body, nil
}
