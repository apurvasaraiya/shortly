package dto

type URLEncodeRequest struct {
	URL string `json:"url"`
}
type URLEncodeResponse struct {
	URL string `json:"url"`
	ID  string `json:"id"`
}
