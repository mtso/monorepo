package server

import (
	"sync"
)

type Store struct {
	sync.Mutex
	mp map[string]string
}

var store Store

func InitStore() {
	store = Store{
		sync.Mutex{},
		make(map[string]string),
	}
}

func (s *Store) Set(key, value string) {
	s.Lock()
	defer s.Unlock()
	s.mp[key] = value
}

func (s *Store) Get(key string) (value string, ok bool) {
	s.Lock()
	defer s.Unlock()
	value, ok = s.mp[key]
	return
}
