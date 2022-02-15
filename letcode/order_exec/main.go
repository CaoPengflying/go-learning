// @Description order_exec
// @Author caopengfei 2021/4/14 20:30
package main

import (
	"fmt"
	"sync"
)

func main() {
	chs := []chan struct{}{make(chan struct{}), make(chan struct{}), make(chan struct{})}

	wg := sync.WaitGroup{}
	wg.Add(3)

	go Cat(chs[0], chs[1], &wg)
	go Dog(chs[1], chs[2], &wg)
	go Pig(chs[2], chs[0], &wg)

	chs[0] <- struct{}{}
	wg.Wait()
}

func Cat(token chan struct{}, nextToken chan struct{}, w *sync.WaitGroup) {
	defer w.Done()

	for i := 0; i < 100; i++ {
		<-token
		fmt.Println("cat", i)
		nextToken <- struct{}{}
	}
	w.Done()
}

func Dog(token chan struct{}, nextToken chan struct{}, w *sync.WaitGroup) {
	defer w.Done()

	for i := 0; i < 100; i++ {
		<-token
		fmt.Println("dog", i)
		nextToken <- struct{}{}
	}
}

func Pig(token chan struct{}, nextToken chan struct{}, w *sync.WaitGroup) {
	defer w.Done()

	for i := 0; i < 100; i++ {
		<-token
		fmt.Println("pig", i)
		if i != 99 {
			nextToken <- struct{}{}
		}
	}
}
