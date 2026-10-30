package service

import (
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
	record := payload.Data[0]
	callingCode := record.CallingCodes[0]
	return "+" + callingCode + " " + phone, nil
}
