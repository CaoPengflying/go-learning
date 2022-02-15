// @Description mutex
// @Author caopengfei 2021/2/7 16:42
package mutex

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type ExtMutex struct {
	m sync.Mutex
}

const (
	mutexLocked = 1 << iota // 1
	mutexWoken              // 2
	mutexStarving
	mutexWaiterShift = iota // 1

)

func (em *ExtMutex) TryLock() bool {
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&em.m)), 0, mutexLocked) {
		return true
	}

	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&em.m)))

	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}

	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&em.m)), old, new)
}

func (em *ExtMutex) Count() int32 {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&em.m)))
	v = v >> mutexWaiterShift
	v = v + (v & mutexLocked)
	return v
}

func (em *ExtMutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&em.m)))
	return state&mutexLocked == mutexLocked
}

func (em *ExtMutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&em.m)))
	return state&mutexWoken == mutexWoken
}

func (em *ExtMutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&em.m)))
	return state&mutexStarving == mutexStarving
}
