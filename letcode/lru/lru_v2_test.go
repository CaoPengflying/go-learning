package lru

import (
	"github.com/stretchr/testify/assert"

	"fmt"
	"testing"
	"time"
)

var lru LRU

func init() {
	lru = NewLru()
}

func TestGetNotFound(t *testing.T) {
	_, err := lru.Get("testKey")

	assert.Equal(t, ErrNotFound, err)
}

func TestFirstSet(t *testing.T) {
	key := "testFirst"
	value := "testValue"

	lru.Set(key, value, 10)

	resValue, err := lru.Get(key)

	assert.Nil(t, err)
	assert.Equal(t, value, resValue)
}

func TestFullSet(t *testing.T) {
	for i := 0; i < 11; i++ {
		lru.Set(fmt.Sprintf("testKey%d", i), fmt.Sprintf("testValue%d", i), -1)
	}

	_, err1 := lru.Get("testKey0")
	value2, err2 := lru.Get("testKey10")

	assert.Equal(t, ErrNotFound, err1)
	assert.Equal(t, nil, err2)
	assert.Equal(t, "testValue10", value2)
}

func TestTimeoutGet(t *testing.T) {
	lru.Set("testKey", "testValue", 1)

	time.Sleep(time.Second * 1)

	_, err := lru.Get("testKey")
	assert.Equal(t, ErrNotFound, err)
}
