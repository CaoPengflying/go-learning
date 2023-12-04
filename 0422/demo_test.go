package main

import (
	"fmt"
	"testing"
)

type Apple struct {
	Name string
}

func (receiver *Apple) name() {
	fmt.Println("hello")
	fmt.Println(receiver.Name)
}

func Test(t *testing.T) {
	var a *Apple
	a.name()
}
