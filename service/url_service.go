package service

import (
	"log"
	"shortly/repository"

	urlpkg "net/url"

	"github.com/google/uuid"
)

var (
	GenerateNewUUIDStringFunc = func() string {
		return uuid.NewString()
	}
)

type URLService interface {
	EncodeURL(url string) (string, error)
	FetchURLFromID(id string) (string, error)
	IncrementVisitCountForHostname(url string) error
	Metrics() (map[string]uint, error)
}

type urlService struct {
	repo repository.Repository
}

func NewURLService(repo repository.Repository) URLService {
	return urlService{repo}
}

// EncodeURL encodes the given url, stores it into database and returns the id
func (s urlService) EncodeURL(url string) (string, error) {
	logger := log.Default()

	id, err := s.repo.FetchIDFromURL(url)
	if err != nil {
		logger.Printf("[error] failed to fetch existing id from url %s %v\n", url, err)
		return "", err
	}

	if id != "" {
		return id, nil
	}

	id = GenerateNewUUIDStringFunc()

	err = s.repo.SaveURLAndId(url, id)
	if err != nil {
		logger.Printf("[error] failed to save url %s with id %s %v", url, id, err)
		return "", err
	}

	return id, nil
}

// FetchURLFromID fetches URL based on given id
func (s urlService) FetchURLFromID(id string) (string, error) {
	logger := log.Default()

	url, err := s.repo.FetchURLFromID(id)
	if err != nil {
		logger.Printf("[error] failed to fetch existing id %s with url %s %v", id, url, err)
		return "", err
	}

	return url, nil
}

// IncrementVisitCountForHostname increments hostname visit by 1.
func (s urlService) IncrementVisitCountForHostname(url string) error {
	urlStruct, err := urlpkg.ParseRequestURI(url)
	if err != nil {
		return err
	}
	return s.repo.IncrementCountForHostname(urlStruct.Hostname())
}

// Metrics fetches top 3 metrics by domain visit
func (s urlService) Metrics() (map[string]uint, error) {
	return s.repo.FetchMetrics(3)
}
