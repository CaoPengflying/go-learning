// @Description fan_in
// @Author caopengfei 2021/7/27 11:28

package main

import (
	"fmt"
	"sync"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				fmt.Println("sq is done")
				return
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, chs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	output := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			select {
			case out <- i:
			case <-done:
				fmt.Println("merge output is done")
				return
			}
			out <- i
		}
	}
	wg.Add(len(chs))
	for _, ch := range chs {
		go output(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := gen(2, 3)
	done := make(chan struct{})

	c1 := sq(done, in)

	c2 := sq(done, in)
	out := merge(done, c1, c2)
	close(done)

	fmt.Println(<-out)
	fmt.Println(<-out)

	time.Sleep(5*time.Second)
}
