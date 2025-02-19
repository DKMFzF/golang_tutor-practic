package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	count int
	mu    *sync.Mutex
}

func (c *counter) inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *counter) value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	counter := counter{
		count: 0,
		mu:    new(sync.Mutex),
	}

	for i := 0; i < 1000; i++ {
		go func() {
			counter.inc()
		}()
	}

	time.Sleep(time.Second)

	fmt.Println(counter.value())
}
