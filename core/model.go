package core

type UrlStore struct {
	URL       string `json:"url" validate:"empty=false & format=url"`
	Visitor   string `json:"visitor"`
	VisitedAt int64  `json:"visited_at,string,omitempty"`
}
