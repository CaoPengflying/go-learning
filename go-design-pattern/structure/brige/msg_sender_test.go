package brige

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorNotification_Notify(t *testing.T) {
	sender := NewTelephoneMsgSender()
	n := ErrorNotification{iMsgSender: sender}
	err := n.Notify("test msg")
	assert.Equal(t, nil, err)
}
