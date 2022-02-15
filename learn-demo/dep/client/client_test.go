package client

import (
	"go-learning/learn-demo/dep/service"
	"testing"
)
//包外访问，不能是单元测试，并且必须是函数首字母大写
func TestDep(t *testing.T) {
	service.Square()
}

