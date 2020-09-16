package _select

import (
	"testing"
	"time"
)

func Service() string {
	time.Sleep(1000 * time.Millisecond)
	return "done"
}

func AsyncService() chan string {
	retCh := make(chan string)

	go func() {
		println("async start")
		ch := Service()
		retCh <- ch
		println("async end")
	}()
	return retCh
}

func Service2() string {
	time.Sleep(100 * time.Millisecond)
	return "done2"
}

func AsyncService2() chan string {
	retCh := make(chan string)

	go func() {
		println("async start")
		ch := Service2()
		retCh <- ch
		println("async end")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case retCh := <-AsyncService():
		t.Log(retCh)
	case <-time.After(time.Millisecond * 500):
		t.Error("timeout")
	case ret := <-AsyncService2():
		t.Log(ret)
	}
}
