package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasPage struct {
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	Count int `json:"count"`
}

func GetLocationAreas(url string) (*LocationAreasPage, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get location areas: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("failed to get location areas: could not read body: %s", err)
	}

	if res.StatusCode >= 300 {
		msg := "failed to get location areas: response failed with status code: %d and\nbody: %s"
		return nil, fmt.Errorf(msg, res.StatusCode, body)
	}

	page := LocationAreasPage{}
	if err := json.Unmarshal(body, &page); err != nil {
		return nil, fmt.Errorf("failed to get location areas: failed to unmarshal body: %s", err)
	}

	return &page, nil
}

func (page *LocationAreasPage) Print() {
	for _, area := range page.Results {
		fmt.Println(area.Name)
	}
}
