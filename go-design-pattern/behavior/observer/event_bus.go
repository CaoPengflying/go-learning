package observer

import (
	"errors"
	"reflect"
	"sync"
)

type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

func (bus *AsyncEventBus) Subscribe(topic string, f interface{}) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	v := reflect.ValueOf(f)

	if v.Type().Kind() != reflect.Func {
		return errors.New("handler is not a function")
	}

	handler, ok := bus.handlers[topic]

	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler, v)
	bus.handlers[topic] = handler

	return nil
}

func (bus *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := bus.handlers[topic]
	if !ok {
		println("not found handlers in topic", topic)
		return
	}

	params := make([]reflect.Value, len(args))

	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for i := range handlers {
		 handlers[i].Call(params)
	}
}
