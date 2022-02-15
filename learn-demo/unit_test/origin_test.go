package unit

import "testing"

//文件必须以_test为结尾
//引入testing包
//测试函数必须是公开的函数，并且以Test开头，参数必须是接收一个指向testing.T类型的指针，并且没有返回值

//mac 快捷键 command + n
func IsEven(num int) bool {
	return num%2 == 0
}

//go 原意使用t.log  t.fatal t.error等来作为测试输出
func TestA(t *testing.T) {
	t.Log("start test Get1")

	expectResult := false
	param := 2

	result := IsEven(param)

	if expectResult == result {
		t.Log("success")
	} else {
		t.Errorf("should receive %v. result is %v", expectResult, result)
	}

}

//表组测试 如果测试接受一组不同的输入产生不同的输出，应该采用表组方式进行测试
func TestTable(t *testing.T) {
	tests := []struct {
		num          int
		expectResult bool
	}{
		{1, false}, {2, true},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := IsEven(test.num)
			if test.expectResult == result {
				t.Log("success")
			} else {
				t.Errorf("should receive %v. result is %v", test.expectResult, result)
			}
		})
	}
}
