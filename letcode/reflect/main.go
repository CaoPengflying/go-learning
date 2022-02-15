// @Description reflect
// @Author caopengfei 2021/4/15 16:59
package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
}

func (a *Animal) Eat() {
	fmt.Println("eat")
}

func (a *Animal) Say(content string) {
	fmt.Println(content)
}

func main() {
	a := Animal{}
	v := reflect.ValueOf(&a)
	m := v.MethodByName("Eat")
	m.Call(make([]reflect.Value, 0))

	m1 := v.MethodByName("Say")
	args := []reflect.Value{reflect.ValueOf("hello")}
	m1.Call(args)
}
