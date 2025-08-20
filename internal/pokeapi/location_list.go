package pokeapi

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type LocationArea struct{
	Name string `json:"name"`
	URL string `json:"url"`
}

type Page struct {
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func (c *Client) GetPageFromReq(pageURL *string) (Page, error) {

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Page{}, fmt.Errorf("error with new req: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Page{}, fmt.Errorf("error getting res: %w", err)
	}

	var page Page
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&page); err != nil {
		return  Page{}, fmt.Errorf("error parsing data: %w", err)
	}
	return page, nil
}

func GetLocations(page Page) []string {
	var locations []string
	for _, result := range page.Results {
		locations = append(locations, result.Name)
	}
	return locations
}

