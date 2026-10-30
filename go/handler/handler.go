package handler

import (
	"encoding/json"
	"net/http"

	"northwind/callbackdirectory/service"
)

type ErrorResponse struct {
	Err string `json:"error"`
}

type SuccessResp struct {
	FmtNumber string
}

type Handler struct {
	svc *service.DirectoryService
}

func NewHandler(svc *service.DirectoryService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) PhoneNumbers(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	phone := r.URL.Query().Get("phone")
	w.Header().Set("Content-Type", "application/json")
	if len(country) == 0 {
		errResp := ErrorResponse{Err: "country is missing"}
		data, _ := json.Marshal(errResp)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(data))
		return
	}
	if len(phone) == 0 {
		errResp := ErrorResponse{Err: "phone is missing"}
		data, _ := json.Marshal(errResp)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(data))
		return
	}
	resolved, err := h.svc.GetPhoneNumbers(country, phone)
	if err != nil {
		if err.Error() == "country not found" {
			errResp := ErrorResponse{Err: resolved}
			data, _ := json.Marshal(errResp)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(data))
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := SuccessResp{FmtNumber: resolved}
	data, _ := json.Marshal(resp)
	w.Write(data)
}

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/phone-numbers", h.PhoneNumbers)
	return mux
}
