package repository

type Repository interface {
	SaveURLAndId(url string, id string) error
	FetchURLIdFromURL(url string) (string, error)
}
