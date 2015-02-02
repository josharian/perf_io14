package set

import (
	"math/rand"
	"testing"
)

func TestBasics(t *testing.T) {
	s := New()
	s.Add(1)
	if s.Contains(0) {
		t.Errorf("contains 0 but 0 has not been added")
	}
	if !s.Contains(1) {
		t.Errorf("does not contain 1 but 1 has been added")
	}
	s.Remove(1)
	if s.Contains(1) {
		t.Errorf("contains 1 but 1 has been removed")
	}
}

func benchmarkSet(b *testing.B, max int, add, contain, remove int) {
	if add+contain+remove != 100 {
		b.Fatalf("add+contain+remove must sum to 100")
	}
	s := New()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			n := rand.Intn(max)
			x := rand.Intn(100)
			switch {
			case x < add:
				s.Add(n)
			case x < add+contain:
				_ = s.Contains(n)
			default:
				s.Remove(n)
			}
		}
	})
}

func BenchmarkSmallReadHeavy(b *testing.B)  { benchmarkSet(b, 10, 10, 80, 10) }
func BenchmarkSmallWriteHeavy(b *testing.B) { benchmarkSet(b, 10, 40, 20, 40) }
func BenchmarkLargeReadHeavy(b *testing.B)  { benchmarkSet(b, 1000, 10, 80, 10) }
func BenchmarkLargeWriteHeavy(b *testing.B) { benchmarkSet(b, 1000, 40, 20, 40) }
func BenchmarkXLReadHeavy(b *testing.B)     { benchmarkSet(b, 10000, 10, 80, 10) }
func BenchmarkXLWriteHeavy(b *testing.B)    { benchmarkSet(b, 10000, 40, 20, 40) }
