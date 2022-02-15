package sync_task

import (
	"runtime"
	"testing"
	"time"
)


func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for i := 0; i < numOfRunner; i++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

// 所有任务执行完成后一起返回，可以是使用waitGroup 同时也可以在用当前方法
func TestAllResponse(t *testing.T) {
	t.Log("before", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("after", runtime.NumGoroutine())
}
