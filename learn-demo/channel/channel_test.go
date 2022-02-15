package channel

import (
	"fmt"
	"strconv"
	"testing"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	for i := 0; i < 10; i++ {
		go func(i int) {
			channel <- strconv.Itoa(i) + "hello"
		}(i)
	}

	msg := <-channel
	t.Log(msg)
}

func TestChannelBuf(t *testing.T) {
	messages := make(chan string, 2)

	messages <- "buffer"
	messages <- "channel"

	t.Log(<-messages)
	t.Log(<-messages)
}

func ping(pings chan string, msg string) {
	pings <- msg
}

func pang(pings chan string, pangs chan<- string) {
	msg := <-pings
	pangs <- msg
}

func TestPingPang(t *testing.T) {
	pings := make(chan string, 1)
	pangs := make(chan string, 1)
	ping(pings, "passed message")
	pang(pings, pangs)
	t.Log(<-pangs)
}

func TestAsyncChannel(t *testing.T) {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		t.Log("receive" + msg)
	default:
		t.Log("receive nothing")

	}

	msg := "hi"
	select {
	case messages <- msg:
		t.Log("send" + msg)
	default:
		t.Log("send nothing")
	}
	select {
	case msg := <-messages:
		t.Log("last receive" + msg)
	case sig := <-signals:
		t.Log(sig)
	default:
		t.Log("no activity")
	}
}

func TestReceive(t *testing.T) {
	ch := make(chan int, 10)
	i := <-ch

	ch <- 1

	fmt.Println(i)
}
