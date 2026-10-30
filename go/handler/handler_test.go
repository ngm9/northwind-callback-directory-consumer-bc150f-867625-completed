package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"northwind/callbackdirectory/client"
	"northwind/callbackdirectory/service"
)

type stubClient struct {
	payload client.CountryPayload
}

func (s stubClient) FetchCountry(country string) (client.CountryPayload, error) {
	return s.payload, nil
}

func TestEndpointKnownCountry(t *testing.T) {
	stub := stubClient{payload: client.CountryPayload{Data: []client.CountryRecord{{Name: "Afghanistan", CallingCodes: []string{"93"}}}}}
	svc := service.NewDirectoryService(stub)
	h := NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/phone-numbers?country=Afghanistan&phone=656445445", nil)
	rec := httptest.NewRecorder()
	h.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}
