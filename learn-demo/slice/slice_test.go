package slice

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	s0 = append(s0, 1, 2, 3)
	t.Log(s0)

	s1 := []int{1, 2, 3}
	t.Log(s1)

	s2 := [2]int{1, 2}
	t.Log(s2)

	s3 := make([]int, 5, 5)
	s3 = append(s3, 1, 2, 3)
	t.Log(s3)

	s4 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1

		s4[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			s4[i][j] = i + j
		}
	}
	fmt.Println(s4)
}

func TestShareMemory(t *testing.T) {
	// cap的计算是从当前指针开始，一直到最后的指针，len为当前切片中元素的个数
	var year = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	t.Log(year, len(year), cap(year))
	//指针在0所以cap是12
	q1 := year[:4]
	t.Log(q1, len(q1), cap(q1))
	//指针在3所以 cap为 12-3 = 9
	q2 := year[3:7]
	t.Log(q2, len(q2), cap(q2))
	q1[0] = "unknow"
	t.Log(year, len(year), cap(year))

}

func TestCompare(t *testing.T) {
	// go 数组可以用==比较，切片不可以
	var s = []int{1, 2, 3}
	var s2 = []int{1, 2, 3}
	//t.Log(s == s2)

	t.Log(s, s2)
	var a = [...]int{1, 2, 3}
	var a2 = [...]int{1, 2, 3}
	t.Log(a == a2)

}

func TestCap(t *testing.T) {
	s1 := make([]int, 5, 8)
	t.Log(len(s1))
	t.Log(cap(s1))
}

func TestExpansion(t *testing.T) {
	slice := []int{1, 2}
	slice = append(slice, 2)
	Param(slice)
	println(slice)
}

func Param(slice []int) {
	//*slice = append(*slice,1)
	slice = append(slice, 1)
	slice[0] = 2
}

func TestAdd(t *testing.T) {
	nums := make([]int, 0, 20)
	add(nums)
	t.Log(nums)
}

func add(nums []int) {
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}
}
