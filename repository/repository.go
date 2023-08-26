package repository

type Repository interface {
	SaveURLAndId(url string, id string) error
	FetchIDFromURL(url string) (string, error)
	FetchURLFromID(id string) (string, error)
}
