package loop

import "testing"

func TestWhileLoop(t *testing.T) {
	//go 中所有的循环都使用 for实现
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
	for {
		t.Log("死循环")
	}
}
