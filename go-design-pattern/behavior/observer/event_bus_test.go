package observer

import (
	"testing"
	"time"
)

func sub1(msg1, msg2 string) {
	time.Sleep(1 * time.Microsecond)
	println("sub1, ", msg1, msg2)
}

func sub2(msg1, msg2 string) {
	println("sub2, ", msg1, msg2)
}

func TestAsyncEventBus_Publish(t *testing.T) {
	bus := NewAsyncEventBus()
	bus.Subscribe("topic:1", sub1)
	bus.Subscribe("topic:1", sub2)
	bus.Publish("topic:1", "test1", "test2")
	bus.Publish("topic:1", "testA", "testB")
	time.Sleep(1 * time.Second)
}
