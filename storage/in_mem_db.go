package storage

import (
	"sync"
)

type inMemoryDB struct {
	m    map[string][]byte
	lock sync.RWMutex
}

// NewInMemoryDB creates a new inMemoryDB
func NewInMemoryDB() DB {
	return &inMemoryDB{m: make(map[string][]byte)}
}

// Get
func (db *inMemoryDB) Get(key string) ([]byte, error) {
	db.lock.RLock()
	defer db.lock.RUnlock()

	val, ok := db.m[key]
	if !ok {
		return nil, ErrNotFound
	}
	return val, nil
}

// Set
func (db *inMemoryDB) Set(key string, val []byte) error {
	db.lock.RLock()
	defer db.lock.RUnlock()
	db.m[key] = val
	return nil
}
