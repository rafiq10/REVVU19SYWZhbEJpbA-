package model

type UrlStore struct {
	URL       string `json:"url"`
	VisitedAt int64  `json:"visited_at"`
}
