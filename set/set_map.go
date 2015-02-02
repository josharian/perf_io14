// +build !ignore

package set

import "sync"

type Set struct {
	mu sync.Mutex
	m  map[int]struct{}
}

func New() *Set {
	return &Set{m: make(map[int]struct{})}
}

func (s *Set) Add(n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[n] = struct{}{}
}

func (s *Set) Contains(n int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.m[n]
	return ok
}

func (s *Set) Remove(n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, n)
}
