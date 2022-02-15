// @Description waitgroup
// @Author caopengfei 2021/2/20 09:39
package main

import (
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()

	wg.Wait()
}
