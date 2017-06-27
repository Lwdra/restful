package storage

import (
	"errors"
)

var (
	// ErrNotFound returned in 404 cases
	ErrNotFound = errors.New("Not found")
)

//DB Key/value storage interface
type DB interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte) error
}
