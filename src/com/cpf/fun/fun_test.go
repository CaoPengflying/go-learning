package fun

import (
	"testing"
	"time"
)

func TestFn(t *testing.T) {
	tsFn := timeSpent(slowFn)
	t.Log(tsFn(10))
}

func slowFn(o int) int {
	time.Sleep(time.Second * 1)
	return o
}

func timeSpent(inner func(o int) int) func(o int) int {
	return func(o int) int {
		start := time.Now()
		ret := inner(o)
		println("time spend ", time.Since(start).Seconds())
		return ret
	}
}

func add(nums ...int) int {
	sum := 0
	for _,val := range nums {
		sum += val
	}
	return sum
}

func TestAdd(t *testing.T) {
	res := add(1,2,3)
	t.Log(res)
}

func TestDefer(t *testing.T) {
	defer t.Log("close resource")
	panic("failed")
}



