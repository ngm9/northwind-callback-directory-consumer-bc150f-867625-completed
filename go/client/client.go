package client

import (
	"encoding/json"
	"net/http"
	"net/url"

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
	url, err := url.Parse(c.cfg.UpstreamBaseURL)
	if err != nil {
		return CountryPayload{}, err
	}
	q := url.Query()
	q.Add("name", country)
	url.RawQuery = q.Encode()
	client := http.Client{
		Timeout: c.cfg.RequestTimeout,
	}
	resp, err := client.Get(url.String())
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
