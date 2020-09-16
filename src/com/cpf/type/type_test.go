package _type

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	//go 不支持隐式类型转换，甚至别名都不允许转换
	var a int32 = 1
	var b int64
	b = int64(a)
	var c = MyInt(b)
	t.Log(c, b, a)

}

func TestPoint(t *testing.T) {
	//go 指针不允许进行算数运算
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
}

func TestString(t *testing.T) {
	// 字符串默认为空字符
	var s string
	t.Log("*" + s + "*")
}
