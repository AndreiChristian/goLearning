package store

import "sync"

type Store struct {
	data map[string]string
	mu   sync.RWMutex
}

func New() *Store {

	return &Store{

		data: make(map[string]string),
	}

}

func (s *Store) Set(key string, value string) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

}

func (s *Store) Get(key string) (string, bool) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.data[key]

	return value, ok

}

func (s *Store) Del(key string) bool {

	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[key]

	if ok {

		delete(s.data, key)

	}

	return ok

}
