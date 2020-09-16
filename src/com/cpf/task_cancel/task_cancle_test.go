package task_cancel

import (
	"testing"
	"time"
)
// 使用广播机制 取消协程
func TestCancelTask(t *testing.T) {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go func(i int) {
			for true {
				if isCanceled(ch) {
					println("canceled", i)
					break
				}
			}
		}(i)
	}
	//cancel(ch)
	closeCancel(ch)
	time.Sleep(1000 * time.Millisecond)
}

func cancel(ch chan int) {
	ch <- 1
}
func closeCancel(ch chan int) {
	close(ch)
}

func isCanceled(ch chan int) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
