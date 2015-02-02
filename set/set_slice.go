// +build ignore

package set

import "sync"

type Set struct {
	mu sync.Mutex
	m  []int
}

func New() *Set {
	return &Set{m: make([]int, 0)}
}

func (s *Set) Add(n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := 0; i < len(s.m); i++ {
		if s.m[i] == n {
			return
		}
	}
	s.m = append(s.m, n)
}

func (s *Set) Contains(n int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := 0; i < len(s.m); i++ {
		if s.m[i] == n {
			return true
		}
	}
	return false
}

func (s *Set) Remove(n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := 0; i < len(s.m); i++ {
		if s.m[i] == n {
			s.m[i] = s.m[len(s.m)-1]
			s.m = s.m[:len(s.m)-1]
			return
		}
	}
}
