package jsn

import (
	"deus-task/core"
	"encoding/json"
	"fmt"

	errs "github.com/pkg/errors"
)

type VisitsSerializer struct{}

func (s *VisitsSerializer) Decode(input []byte) (*core.UrlStore, error) {
	store := &core.UrlStore{}
	if err := json.Unmarshal(input, store); err != nil {
		fmt.Println(input, err)
		return nil, errs.Wrap(err, "serializer.VisitsSerializer.Decode()")
	}
	return store, nil
}

func (s *VisitsSerializer) Encode(input *core.UrlStore) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errs.Wrap(err, "serializer.VisitsSerializer.Encode()")
	}
	return rawMsg, nil
}
