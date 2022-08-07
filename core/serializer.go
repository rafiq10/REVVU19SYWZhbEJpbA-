package core

type VisitsSerializer interface {
	Decode(input []byte) (*UrlStore, error)
	Encode(input *UrlStore) ([]byte, error)
}
