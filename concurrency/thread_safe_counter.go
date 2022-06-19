package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu   sync.RWMutex
	data map[string]int
}

func (c *SafeCounter) Increment(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

func usingSafeCounter() {
	c := SafeCounter{data: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Increment("key")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("key"))
}
