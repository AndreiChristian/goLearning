package store

import "sync"

type Store struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func New() *Store {

	return &Store{

		data: make(map[string]interface{}),
	}

}

func (s *Store) Set(key string, value string) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

}

func (s *Store) Get(key string) (interface{}, bool) {

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

func (s *Store) LPUSH(key string, values ...string) {

	s.mu.Lock()
	defer s.mu.Unlock()

	list, ok := s.data[key].([]string)

	if !ok {
		list = make([]string, 0)
	}

	list = append(list, values...)

}

func (s *Store) RPop(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	list, ok := s.data[key].([]string)
	if !ok || len(list) == 0 {
		return "", false
	}

	lastElem := list[len(list)-1]
	list = list[:len(list)-1]

	s.data[key] = list
	return lastElem, true
}
