package client

import (
	"encoding/json"
	"net/http"

	"northwind/callbackdirectory/config"
)

type CountryRecord struct {
	Name         string   `json:"name"`
	CallingCodes []string `json:"callingCodes"`
}

type CountryPayload struct {
	Data []CountryRecord `json:"data"`
}

type CountriesClient interface {
	FetchCountry(country string) (CountryPayload, error)
}

type HTTPCountriesClient struct {
	cfg config.Config
}

func NewHTTPCountriesClient(cfg config.Config) *HTTPCountriesClient {
	return &HTTPCountriesClient{cfg: cfg}
}

func (c *HTTPCountriesClient) FetchCountry(country string) (CountryPayload, error) {
	url := "https://jsonmock.hackerrank.com/api/countries?name=" + country
	resp, err := http.Get(url)
	if err != nil {
		return CountryPayload{}, err
	}
	defer resp.Body.Close()

	var payload CountryPayload
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return CountryPayload{}, err
	}
	return payload, nil
}
