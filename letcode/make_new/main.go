// @Description make_new
// @Author caopengfei 2021/4/20 10:23
package main

import "fmt"

type A struct {
	name string
	age int
}

func (a *A) Hello() {
	fmt.Println("hello")
}

func main() {
	a := new(A)
	a.Hello()
	fmt.Println(a)



}
