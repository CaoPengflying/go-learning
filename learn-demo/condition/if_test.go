package condition

import "testing"

func TestIf(t *testing.T) {
	//go 可以书写成二段式，可以缩小变量的作用域
	if val, err := fun1(); err == nil {
		t.Log(val)
	}
}

func fun1() (string, error) {
	return "ok", nil
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2, 4:
			t.Log("偶数")
		case 1, 3:
			t.Log("奇数")
		default:
			t.Log("不存在")
		}

	}
}

func TestConditionSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		/*switch i%2 {
		case 0:
			t.Log("偶数")
		case 1:
			t.Log("奇数")
		default:
			t.Log("不存在")
		}*/
		switch {
		case i%2 == 0:
			t.Log("偶数")
		case i%2 == 1:
			t.Log("奇数")
		default:
			t.Log("不存在")
		}

	}
}
