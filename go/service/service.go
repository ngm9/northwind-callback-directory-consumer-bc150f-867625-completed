package service

import (
	"fmt"
	"northwind/callbackdirectory/client"
)

type DirectoryService struct {
	client client.CountriesClient
}

func NewDirectoryService(c client.CountriesClient) *DirectoryService {
	return &DirectoryService{client: c}
}

func (s *DirectoryService) GetPhoneNumbers(country string, phone string) (string, error) {
	payload, err := s.client.FetchCountry(country)
	if err != nil {
		return "", err
	}
	if len(payload.Data) == 0 {
		return "-1", fmt.Errorf("country not found")
	}
	record := payload.Data[0]
	callingCode := record.CallingCodes[len(record.CallingCodes)-1]
	return "+" + callingCode + " " + phone, nil
}
