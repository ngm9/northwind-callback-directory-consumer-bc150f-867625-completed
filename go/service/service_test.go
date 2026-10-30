package service

import (
	"testing"

	"northwind/callbackdirectory/client"
)

type stubClient struct {
	payload client.CountryPayload
}

func (s stubClient) FetchCountry(country string) (client.CountryPayload, error) {
	return s.payload, nil
}

func TestSingleCallingCode(t *testing.T) {
	stub := stubClient{payload: client.CountryPayload{Data: []client.CountryRecord{{Name: "Afghanistan", CallingCodes: []string{"93"}}}}}
	svc := NewDirectoryService(stub)
	result, err := svc.GetPhoneNumbers("Afghanistan", "656445445")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "+93 656445445" {
		t.Fatalf("expected +93 656445445, got %s", result)
	}
}
