package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/encode/", encodeHandler)
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}

type encodeRequest struct {
	URL string `json:"url"`
}

func encodeHandler(w http.ResponseWriter, r *http.Request) {
	logger := log.Default()
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.Printf("[error] invalid method %s\n", r.Method)
		return
	}

	var body encodeRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Printf("[error] error decoding request %v\n", err)
		return
	}
	defer r.Body.Close()

	logger.Printf("body>>>>>>>>>>> %v\n", body)
	w.WriteHeader(http.StatusOK)
}
