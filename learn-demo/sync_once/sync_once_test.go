package sync_once

import (
	"testing"
	"time"
)

func TestFun(t *testing.T) {
	for i := 0; i < 3; i++ {
		go func() {
			Fun()
		}()
	}
	time.Sleep(1 * time.Second)
}
