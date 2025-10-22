package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	v  map[string]int
	mu sync.RWMutex
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func (c *SafeCounter) Print() {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println(c.v)
}

func (c *SafeCounter) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.v, key)
}

func main() {
	sc := SafeCounter{v: make(map[string]int)}

	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		defer wg.Done()
		sc.Inc("key1")
	}()

	go func() {
		defer wg.Done()
		sc.Value("key1")
	}()

	go func() {
		defer wg.Done()
		sc.Delete("key1")
	}()

	go func() {
		defer wg.Done()
		sc.Print()
	}()

	wg.Wait()
}
