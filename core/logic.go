package core

import (
	"time"

	errs "github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
	// "time"
)

type VisitService struct {
	urlRepo VisitsRepository
}

func NewVisitService(repo VisitsRepository) *VisitService {
	return &VisitService{urlRepo: repo}
}

func (srv *VisitService) GetUniqueVisitsNumber(url string) (numVisits int, err error) {
	return srv.urlRepo.GetUniqueVisitsNumber(url)
}

func (srv *VisitService) SaveVisit(urlData *UrlStore) error {
	if err := validate.Validate(urlData); err != nil {
		return errs.Wrap(ErrUrlInvalid, "service.VisitService.SaveVisit()")
	}
	urlData.VisitedAt = time.Now().Unix()
	return srv.urlRepo.SaveVisit(urlData)
}
