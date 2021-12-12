package set

import (
	"sync"
)

type Set interface {
	Add(interface{}) bool
	AddMany(...interface{})
	Remove(interface{}) bool
	Has(interface{}) bool
	Len() int
	Slice() []interface{}
}

type void struct{}

type set struct {
	sync.Mutex
	m map[interface{}]void
}

func NewSet(items ...interface{}) Set {
	set := &set{
		m: make(map[interface{}]void),
	}

	set.AddMany(items...)

	return set
}

func (s *set) Add(i interface{}) bool {
	s.Lock()
	defer s.Unlock()

	if _, found := s.m[i]; found {
		return false
	}

	s.m[i] = void{}

	return true
}

func (s *set) AddMany(items ...interface{}) {
	s.Lock()
	defer s.Unlock()

	for _, i := range items {
		s.m[i] = void{}
	}
}

func (s *set) Remove(i interface{}) bool {
	s.Lock()
	defer s.Unlock()

	if _, found := s.m[i]; found {
		return false
	}

	delete(s.m, i)

	return true
}

func (s *set) Has(i interface{}) bool {
	s.Lock()
	defer s.Unlock()

	_, has := s.m[i]

	return has
}

func (s *set) Len() int {
	s.Lock()
	defer s.Unlock()

	return len(s.m)
}

func (s *set) Slice() []interface{} {
	s.Lock()
	defer s.Unlock()

	k := make([]interface{}, len(s.m))
	for key := range s.m {
		k = append(k, key)
	}

	return k
}
