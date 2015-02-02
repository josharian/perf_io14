package counter

import (
	"math/rand"
	"sync"
)

const shards = 1

type Counter struct {
	mu [shards]sync.Mutex
	n  [shards]int
}

func (c *Counter) Increment() {
	r := rand.Intn(shards)
	c.mu[r].Lock()
	c.n[r]++
	c.mu[r].Unlock()
}

func (c *Counter) Value() int {
	v := 0
	for i := 0; i < shards; i++ {
		c.mu[i].Lock()
		defer c.mu[i].Unlock()
		v += c.n[i]
	}
	return v
}
