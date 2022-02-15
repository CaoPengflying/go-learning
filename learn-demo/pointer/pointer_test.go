package pointer

import "testing"

func zeroVal(val int)  {
	val = 0
}

func zeroPtr(valPtr *int) {
	*valPtr = 0
}

func TestZero(t *testing.T) {
	i:=1
	t.Log(&i)
	t.Log(i)
	zeroVal(i)
	t.Log(i)
	zeroPtr(&i)
	t.Log(i)
	t.Log(&i)

}

