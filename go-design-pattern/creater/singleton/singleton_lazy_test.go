package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLazyInstance(t *testing.T) {
	s := singleton
	t.Log(s)
	ls := lazySingleton
	t.Log(ls)
	s1 := GetLazyInstance()
	s2 := GetLazyInstance()

	assert.Equal(t, s1, s2)
}
