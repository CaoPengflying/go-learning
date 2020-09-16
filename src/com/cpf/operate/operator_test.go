package operate

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	//b := [...]int{1, 2, 3, 4, 5}
	c := [...]int{1, 2, 4, 5}
	d := [...]int{1, 2, 3, 4}
	//长度不同不允许比较
	//t.Log(a == b)
	t.Log(a == c)
	t.Log(a == d)
}

const  (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClea(t *testing.T) {
	// 0111
	a := 7
	// 0111 &^ 1 0110
	a = a &^ Readable
	// 0110 &^ 0100 0010
	a = a &^ Executable
	t.Log(a)
}

