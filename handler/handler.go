package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"shortly/dto"
	"shortly/helper"
	"shortly/service"
)

type URLHandler interface {
	EncodeURL(http.ResponseWriter, *http.Request)
	Redirect(http.ResponseWriter, *http.Request)
	Metrics(http.ResponseWriter, *http.Request)
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

	_, err = url.ParseRequestURI(req.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		helper.WriteError(w, "invalid url %s", req.URL)
		logger.Printf("%v\n", err)
		return
	}

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

// Redirect handles all incoming request for path [GET] /
func (h urlHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	logger := log.Default()

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.Printf("[error] invalid request method %s\n", r.Method)
		return
	}

	id := r.URL.EscapedPath()[len("/"):]
	url, err := h.urlService.FetchURLFromID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Printf("[error] failed to fetch url for id %s\n", id)
		return
	}

	if url == "" {
		w.WriteHeader(http.StatusNotFound)
		helper.WriteError(w, "url not found for id %s", id)
		return
	}

	//ignoring error as updating metrics should not affect redirect flow
	h.urlService.IncrementVisitCountForHostname(url)

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

// Metrics handles all incoming request for path [GET] /metrics
func (h urlHandler) Metrics(w http.ResponseWriter, r *http.Request) {
	logger := log.Default()

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.Printf("[error] invalid request method %s\n", r.Method)
		return
	}

	metrics, err := h.urlService.Metrics()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Printf("[error] failed to fetch metrics %v\n", err)
		return
	}

	err = json.NewEncoder(w).Encode(metrics)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Printf("[error] failed to encode metrics to json %v\n", err)
		return
	}

}
