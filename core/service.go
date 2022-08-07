package core

import "errors"

var (
	ErrUrlNotFound = errors.New("URL not found")
	ErrUrlInvalid  = errors.New("URL not valid")
)

type VisitsService interface {
	GetUniqueVisitsNumber(url string) (numVisits int, err error)
	SaveVisit(urlData *UrlStore) (err error)
}
