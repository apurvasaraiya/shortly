package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"shortly/dto"
	"shortly/service"
)

type URLHandler interface {
	EncodeURL(http.ResponseWriter, *http.Request)
}

type urlHandler struct {
	urlService service.URLService
}

func NewHandler(urlService service.URLService) URLHandler {
	return urlHandler{urlService}
}

// EncodeURL handles all incoming request to path [POST] /encode/
func (h urlHandler) EncodeURL(w http.ResponseWriter, r *http.Request) {
	logger := log.Default()

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.Printf("[error] invalid request method %s\n", r.Method)
		return
	}

	var req dto.URLEncodeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Printf("[error] failed to decode request body %v\n", err)
		return
	}
	defer r.Body.Close()

	id, err := h.urlService.EncodeURL(req.URL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Printf("[error] failed to encode request url %v\n", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	resp := dto.URLEncodeResponse{URL: req.URL, ID: id}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Printf("[error] failed to encode response body %v\n", err)
		return
	}
}
