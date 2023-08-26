package inmemory

import (
	"shortly/repository"
)

type inmemoryDB struct {
	urlToID map[string]string
	idToUrl map[string]string
}

func NewInMemoryDB() repository.Repository {
	return inmemoryDB{
		urlToID: make(map[string]string),
		idToUrl: make(map[string]string),
	}
}

// FetchIDFromURL fetches ID from given url
func (db inmemoryDB) FetchIDFromURL(url string) (string, error) {
	return db.urlToID[url], nil
}

// SaveURLAndId stores url and id
func (db inmemoryDB) SaveURLAndId(url string, id string) error {
	db.urlToID[url] = id
	db.idToUrl[id] = url

	return nil
}

// FetchURLFromID fetches URL for given id
func (db inmemoryDB) FetchURLFromID(id string) (string, error) {
	return db.idToUrl[id], nil
}
