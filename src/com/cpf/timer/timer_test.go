package timer

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	t.Log("timer 1 expired")
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		t.Log("timer 2 expired")
	}()
	time.Sleep(time.Second)
	stop2 := timer2.Stop()
	if stop2 {
		t.Log("timer 2 stopped")
	} else {
		t.Log("timer 2  expired and cannot stopped")
	}
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t2 := range ticker.C {
			t.Log("tick at ", t2)
		}
	}()
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	t.Log("ticker stopped")
}

func TestFrequency(t *testing.T) {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 1000)

	for request := range requests {
		<-limiter
		t.Log("request", request, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	go func() {
		for t := range time.Tick(time.Millisecond * 1000) {
			burstyLimiter <- t
		}
	}()
	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for request := range burstyRequests {
		<-burstyLimiter
		t.Log("request", request, time.Now())
	}
}
