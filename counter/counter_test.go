package counter

import "testing"

func BenchmarkCounter(b *testing.B) {
	var c Counter
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Increment()
		}
	})

	if val := c.Value(); val != b.N {
		b.Errorf("value: got %d want %d", val, b.N)
	}
}
