// @Description _defer
// @Author caopengfei 2021/4/28 19:31
package main

func main() {
	p, b := foo()
	println(p.Name, b)
}

type Person struct {
	Name string
}

func foo() (*Person, int) {
	a := Person{Name: "a"}
	b := 1
	defer func() {
		a.Name = "b"
		b = 2
	}()
	return &a, b
}
