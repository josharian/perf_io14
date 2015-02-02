// +build ignore

package set

import "sync"

type Set struct {
	mu sync.Mutex
	m  map[int]int
}

func New() *Set {
	return &Set{m: make(map[int]int)}
}

func (s *Set) Add(n int) {
	s.mu.Lock()
	s.m[n]++
	s.mu.Unlock()
}

func (s *Set) Contains(n int) bool {
	s.mu.Lock()
	i := s.m[n]
	ok := i > 0
	s.mu.Unlock()
	return ok
}

func (s *Set) Remove(n int) {
	s.mu.Lock()
	s.m[n]--
	s.mu.Unlock()
}
