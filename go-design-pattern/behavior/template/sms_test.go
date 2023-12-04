package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTelecomSms_Send(t *testing.T) {
	tel := NewTelecomSms()
	err := tel.Send("test", "10086")
	assert.NoError(t, err)
}
