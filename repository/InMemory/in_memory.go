package inmemory

import (
	"deus-task/core"
	"sync"
	"time"

	errs "github.com/pkg/errors"
)

type InMemoryStore struct {
	mu       sync.Mutex
	client   map[string]int
	database string
	timeout  time.Duration
}

func NewInMemoryStore(db string, timeout int) (core.VisitsRepository, error) {
	client := make(map[string]int)
	return &InMemoryStore{client: client, database: db, timeout: time.Duration(timeout) * time.Second}, nil
}
func (s *InMemoryStore) GetUniqueVisitsNumber(url string) (numVisits int, err error) {
	numVisits, ok := s.client[url]
	if !ok {
		return 0, errs.Wrap(core.ErrUrlNotFound, "repository.InMemoryStore.GetUniqueVisitsNumber()")
	}
	return numVisits, nil
}

func (s *InMemoryStore) SaveVisit(urlData *core.UrlStore) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.client[urlData.URL]; ok {
		s.client[urlData.URL]++
	} else {
		s.client[urlData.URL] = 1
	}
	return nil
}
