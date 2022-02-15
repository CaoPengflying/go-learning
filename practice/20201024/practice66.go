package main

import "fmt"
//fmt.print函数调用的是结构体的String()方法
//%#表示打印结构体信息
type T struct {
	x int
}

func (t T) String() string {
	return "boo"
}

func main() {
	t := T{123}
	fmt.Printf("%v\n",t)
	fmt.Printf("%#v\n",t)
}
