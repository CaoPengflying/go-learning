package proxy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	proxy := UserControllerProxy{}
	result := proxy.login("cpf", "123456")
	assert.Equal(t, true, result)
}
