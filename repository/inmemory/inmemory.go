package inmemory

import (
	"shortly/helper"
	"shortly/repository"
	"sync"
)

type inmemoryDB struct {
	mu               sync.Mutex
	urlToID          map[string]string
	idToUrl          map[string]string
	domainVisitCount map[string]uint
}

func NewInMemoryDB() repository.Repository {
	return &inmemoryDB{
		urlToID:          make(map[string]string),
		idToUrl:          make(map[string]string),
		domainVisitCount: make(map[string]uint),
	}
}

// FetchIDFromURL fetches ID from given url
func (db *inmemoryDB) FetchIDFromURL(url string) (string, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	return db.urlToID[url], nil
}

// SaveURLAndId stores url and id
func (db *inmemoryDB) SaveURLAndId(url string, id string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.urlToID[url] = id
	db.idToUrl[id] = url

	return nil
}

// FetchURLFromID fetches URL for given id
func (db *inmemoryDB) FetchURLFromID(id string) (string, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	return db.idToUrl[id], nil
}

// IncrementCountForHostname increments visit count for given url
func (db *inmemoryDB) IncrementCountForHostname(url string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.domainVisitCount[url] = db.domainVisitCount[url] + 1
	return nil
}

// FetchMetrics fetches metrics for all urls
func (db *inmemoryDB) FetchMetrics(topN int) (map[string]uint, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	return helper.SortTopNInMap(db.domainVisitCount, 3), nil
}
