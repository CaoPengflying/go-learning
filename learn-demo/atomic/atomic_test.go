package atomic

import (
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	var ops uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for true {
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)

	t.Log("ops:", opsFinal)

}
