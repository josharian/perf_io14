// +build ignore

package set

import (
	"sort"
	"sync"
)

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
	i := sort.SearchInts(s.m, n)
	if i == len(s.m) {
		s.m = append(s.m, n)
		return
	}
	if s.m[i] == n {
		return
	}
	s.m = append(s.m, 0)
	copy(s.m[i:], s.m[i+1:])
	s.m[i] = n
}

func (s *Set) Contains(n int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	i := sort.SearchInts(s.m, n)
	return i < len(s.m) && s.m[i] == n
}

func (s *Set) Remove(n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	i := sort.SearchInts(s.m, n)
	if i == len(s.m) || s.m[i] != n {
		return
	}
	copy(s.m[i+1:], s.m[i:])
	s.m = s.m[:len(s.m)-1]
}
