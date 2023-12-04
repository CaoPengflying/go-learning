package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClone(t *testing.T) {
	updateAt, _ := time.Parse("2006", "2020")
	words := KeyWords{
		"testA": &KeyWord{
			word:     "testA",
			visit:    1,
			UpdateAt: &updateAt,
		},
		"testB": &KeyWord{
			word:     "testB",
			visit:    2,
			UpdateAt: &updateAt,
		},
	}
	now := time.Now()
	updateWords := []*KeyWord{
		{
			word:     "testB",
			visit:    10,
			UpdateAt: &now,
		},
	}
	got := words.Clone(updateWords)

	assert.Equal(t, words["testA"], got["testA"])
	assert.Equal(t, words["testB"], got["testB"])
	assert.Equal(t, updateWords[0], got["testB"])
	assert.Equal(t, words["testC"], got["testC"])
}
