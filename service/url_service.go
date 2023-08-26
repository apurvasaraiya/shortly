package service

import (
	"log"
	"shortly/repository"

	"github.com/google/uuid"
)

type URLService interface {
	EncodeURL(url string) (string, error)
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

	id, err := s.repo.FetchURLIdFromURL(url)
	if err != nil {
		logger.Printf("[error] failed to fetch existing id from url %s %v\n", url, err)
		return "", err
	}

	if id != "" {
		return id, nil
	}

	id = uuid.NewString()

	err = s.repo.SaveURLAndId(url, id)
	if err != nil {
		logger.Printf("[error] failed to save url %s with id %s %v", url, id, err)
		return "", err
	}

	return id, nil
}
