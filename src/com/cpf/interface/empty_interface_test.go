package _interface

import "testing"

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		println("int", i)
	}else if i, ok := p.(string); ok {
		println("string", i)
	}else {
		println("unKnown")
	}

	switch v := p.(type) {
	case int:
		println("switch int", v)
	case string:
		println("switch string", v)
	default:
		println("switch unKnown")
	}
}

func TestEmptyInterface(t *testing.T) {
	a := 1
	b := "hello"
	DoSomething(a)
	DoSomething(b)
	c := false
	DoSomething(c)
}
