package _map

import "testing"

func TestMapInit(t *testing.T) {
	// go 中map key不存在时,val也会赋值为0，判断key是否存在，map[]方法有两个返回值，第二个返回值为一个布尔类型
	var m1 = map[int]int{}

	m1[1] = 1
	t.Log(m1)

	var m2 = map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m2)

	if val, ok := m2[4]; ok {
		t.Log("value is ", val)
	} else {
		t.Log("key is not exist")
	}

}

func TestMapTraverse(t *testing.T) {
	m1 := make(map[int]int)
	m1[0] = 1
	m1[2] = 3
	for i := range m1 {
		t.Log(i)
	}
	for key, val := range m1 {
		t.Log(key,val)
	}
}

func TestMapFun(t *testing.T) {
	mf := map[int]func(op int)int{}
	mf[1] = func(op int) int {
		return op
	}
	mf[2] = func(op int) int {
		return op * op
	}
	mf[3] = func(op int) int {
		return op * op * op
	}

	t.Log(mf[1](2))
	t.Log(mf[2](2))
	t.Log(mf[3](2))
}

