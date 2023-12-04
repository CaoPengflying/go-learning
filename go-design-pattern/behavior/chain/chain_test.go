package chain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterChain_Filter(t *testing.T) {
	chain := FilterChain{}
	chain.AddFilter(&SexWordFilter{})
	chain.AddFilter(&PoliticalWordFilter{})

	res := chain.Filter("发布内容")
	assert.True(t, res)
}
