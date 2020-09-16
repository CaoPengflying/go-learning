package array

import (
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr1 [3]int
	arr1[0] = 1
	t.Log(arr1)

	arr2 := [3]int{1, 2, 3}
	t.Log(arr2)
	arr3 := [...]int{1, 2, 3, 4}
	t.Log(arr3)

	arr4 := [2][2]int{{1, 2}, {3, 4}}
	t.Log(arr4)
}

func TestArrayTraverse(t *testing.T) {

	arr := [...]int{1, 2, 3, 4}

	for i, e := range arr {
		t.Log(i, e)
	}
	for i := range arr {
		t.Log(i)
	}
	for i := 0; i < len(arr); i++ {
		t.Log(i, arr[i])
	}
}

func TestSlice(t *testing.T) {
	//go 支持数组截取
	arr := [4]int{1, 2, 3, 4}
	arr1 := arr[:3]
	t.Log(arr)
	t.Log(arr1)
	arr2 := arr[1:3]
	t.Log(arr2)
}
