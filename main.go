package main

import (
	"log"
	"net/http"
	"shortly/handler"
	"shortly/repository/inmemory"
	"shortly/service"
)

func main() {
	repo := inmemory.NewInMemoryDB()
	s := service.NewURLService(repo)
	h := handler.NewHandler(s)
	http.HandleFunc("/encode/", h.EncodeURL)
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}
