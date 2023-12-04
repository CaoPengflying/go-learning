package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				println(err)
			}
		}()
		panic("panic go1")
	}()
	/*	go func() {
		panic("panic go2")
	}()*/

	panic("panic")

	fmt.Println("hello world")
	time.Sleep(1 * time.Second)
}
