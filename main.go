package main

import (
	"log"
	"net/http"
	"shortly/handler"
)

func main() {
	h := handler.NewHandler()
	http.HandleFunc("/encode/", h.EncodeURL)
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}
