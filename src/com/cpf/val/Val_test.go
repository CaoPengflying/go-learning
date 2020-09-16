package val

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int = 1
	//var (
	//	a = 1
	//	b = 1
	//)
	// go是一个强类型语言，但是go有类型推断功能，以下是最简单的赋值方式
	a := 1
	b := 1
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println(" ", 1/3)
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	println("change before", a, b)
	/*
		tmp := a
		a = b
		b = tmp*/
	a, b = b, a
	println("change after", a, b)
}

//连续常量的声明
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday)
	t.Log(Readable, Writable, Executable)
}
