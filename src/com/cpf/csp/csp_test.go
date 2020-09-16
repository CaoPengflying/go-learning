package csp

import (
	"testing"
	"time"
)

func serviceTask() string {
	time.Sleep(500 * time.Millisecond)
	return "service task end"
}

func otherTask() {
	println("work other things")
	time.Sleep(1000 * time.Millisecond)
	println("other task end")
}

func TestService(t *testing.T) {
	println(serviceTask())
	otherTask()
}

func AsyncService() chan string {
	//retCh := make(chan string)
	retCh := make(chan string,1)
	go func() {
		ret := serviceTask()
		println("returned result .")
		retCh <- ret
		println("service exit")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	println(<-retCh)
}
