package handler

import (
	"net/http"

	"northwind/callbackdirectory/service"
)

type Handler struct {
	svc *service.DirectoryService
}

func NewHandler(svc *service.DirectoryService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) PhoneNumbers(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	phone := r.URL.Query().Get("phone")
	resolved, err := h.svc.GetPhoneNumbers(country, phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(resolved))
}

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/phone-numbers", h.PhoneNumbers)
	return mux
}
