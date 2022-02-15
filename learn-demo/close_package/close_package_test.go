package close_package

import "testing"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func TestIntSeq(t *testing.T) {
	nextInt := intSeq()
	t.Log(nextInt())
	t.Log(nextInt())
	t.Log(nextInt())
	newInts := intSeq()
	t.Log(newInts())
}
