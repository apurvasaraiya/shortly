package repository

type Repository interface {
	SaveURLAndId(url string, id string) error
	FetchIDFromURL(url string) (string, error)
	FetchURLFromID(id string) (string, error)
	IncrementCountForHostname(hostname string) error
	FetchMetrics(topN int) (map[string]uint, error)
}
