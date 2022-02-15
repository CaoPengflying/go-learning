// @Description rwmutex
// @Author caopengfei 2021/2/18 18:00
package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.RWMutex
	count int
}

func (c *Counter) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.count
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func foo(rwm *sync.RWMutex) {
	fmt.Println("in foo")
	rwm.RLock()
	bar(rwm)
	rwm.RUnlock()
}
func bar(rwm *sync.RWMutex) {
	rwm.Lock()
	fmt.Println("in bar")
	rwm.Unlock()
}

func main() {
	var m sync.RWMutex
	go func() {
		time.Sleep(200 * time.Millisecond)
		m.Lock()
		fmt.Println("Lock")
		time.Sleep(100 * time.Millisecond)
		m.Unlock()
		fmt.Println("Unlock")
	}()

	factorial(&m, 10)
}

func factorial(m *sync.RWMutex, n int) int {
	if n < 1 {
		return 0
	}
	fmt.Println("RLock")
	m.RLock()
	defer func() {
		fmt.Println("RUnlock")
		m.RUnlock()
	}()

	time.Sleep(100 * time.Millisecond)
	return factorial(m, n-1) * n
}
