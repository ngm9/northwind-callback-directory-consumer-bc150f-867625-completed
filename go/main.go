package main

import (
	"log"
	"net/http"

	"northwind/callbackdirectory/client"
	"northwind/callbackdirectory/config"
	"northwind/callbackdirectory/handler"
	"northwind/callbackdirectory/service"
)

func main() {
	cfg := config.Load()
	c := client.NewHTTPCountriesClient(cfg)
	svc := service.NewDirectoryService(c)
	h := handler.NewHandler(svc)
	log.Println("Callback directory listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", h.Routes()))
}
