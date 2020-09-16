package sync_task

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)

	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}
// 多个请求，谁先响应就返回谁的内容
func TestFirstResponse(t *testing.T) {
	t.Log("before", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("after",runtime.NumGoroutine())
}
