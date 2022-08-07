package core

type VisitsRepository interface {
	GetUniqueVisitsNumber(url string) (numVisits int, err error)
	SaveVisit(urlData *UrlStore) (err error)
}
