package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type URLHandler interface {
	EncodeURL(http.ResponseWriter, *http.Request)
}

type urlHandler struct{}

func NewHandler() URLHandler {
	return urlHandler{}
}

// EncodeURL handles all incoming request to path [POST] /encode/
func (h urlHandler) EncodeURL(w http.ResponseWriter, r *http.Request) {
	logger := log.Default()

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.Printf("[error] invalid request method %s\n", r.Method)
		return
	}

	var req struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Printf("[error] failed to decode request body %v\n", err)
		return
	}
	defer r.Body.Close()

	logger.Println("req>>>", req)
	w.WriteHeader(http.StatusOK)
}
