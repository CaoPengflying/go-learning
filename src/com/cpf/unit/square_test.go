package unit_test

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"go-learning/src/com/cpf/unit"
	"testing"
)
//go test -v -cover 单元测试覆盖率
func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expects := [...]int{1, 4, 9}

	for i, v := range inputs {
		assert.Equal(t, unit.Square(v), expects[i])
	}
}
// 单元测试中 error会停止执行该方法，继续执行其他的方法
// fatal 继续执行该方法
func TestAdd2(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	//expects := [...]int{1, 4, 9}

	for _, v := range inputs {
		actural := unit.Add(v,v)
		println(actural)
		t.Error("error")
	}

}
func TestAdd(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	for _, v := range inputs {
		actural := unit.Add(v,v)
		println(actural)
		t.Fatal("error")
	}
}



