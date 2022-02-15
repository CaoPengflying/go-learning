package _interface

import (
	"testing"
	"time"
)

// go 中的接口实现，不需要依赖，是非入侵性的
type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "go hello world"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Logf(p.WriteHelloWorld())
}

type IntConv func(o int) int

func timeSpent(inner IntConv) IntConv {
	return func(o int) int {
		start := time.Now()
		ret := inner(o)
		println("time spend ", time.Since(start).Seconds())
		return ret
	}
}
func slowFn(o int) int {
	time.Sleep(time.Second * 1)
	return o
}
func TestFn(t *testing.T) {
	tsFn := timeSpent(slowFn)
	t.Log(tsFn(10))
}

type JavaProgrammer struct {

}

func (g *JavaProgrammer) WriteHelloWorld() string {
	return "java hello world"
}

func writeFirstProgrammer(p Programmer)  {
	println(p.WriteHelloWorld())
}
//多态
func TestPolymorphism(t *testing.T) {
	g := new(GoProgrammer)
	j := new(JavaProgrammer)
	writeFirstProgrammer(g)
	writeFirstProgrammer(j)
}
