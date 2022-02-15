// @Description main
// @Author caopengfei 2021/2/5 13:49
package mutex

import "sync/atomic"

type MyMutex struct {
	state int32
	sema  uint32
}



func (m *MyMutex) Lock() {
	// fast path : 幸运case， 能够直接获取到锁
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}

	awoke := false
	for {
		old := m.state
		new := old | mutexLocked

		if old&mutexLocked != 0 {
			new = old + 1<<mutexWaiterShift //等待着数量+1
		}

		if awoke {
			// goroutine 是唤醒的 0011  0001
			new &^= mutexWoken
		}

		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			if old&mutexLocked == 0 {
				break
			}
			//runtime.Semacquire(&m.sema)
			awoke = true
		}
	}
}
