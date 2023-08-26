package inmemory

import (
	"shortly/repository"
)

type inmemoryDB struct {
	urlToID map[string]string
	IDToUrl map[string]string
}

func NewInMemoryDB() repository.Repository {
	return inmemoryDB{
		urlToID: make(map[string]string),
		IDToUrl: make(map[string]string),
	}
}

// FetchURLIdFromURL fetches URL id from given url
func (db inmemoryDB) FetchURLIdFromURL(url string) (string, error) {
	var id string
	var ok bool

	if id, ok = db.urlToID[url]; !ok {
		return "", nil
	}

	return id, nil
}

// SaveURLAndId stores url and id
func (db inmemoryDB) SaveURLAndId(url string, id string) error {
	db.urlToID[url] = id
	db.IDToUrl[id] = url

	return nil
}
