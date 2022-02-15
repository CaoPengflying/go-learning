// @Description mutex
// @Author caopengfei 2021/2/7 16:42
package mutex

import (
	"fmt"
	"github.com/petermattis/goid"
	"sync"
	"sync/atomic"
)

type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // 当前持有锁的goroutine id
	recursion int32 // 重入的次数
}

func (rm *RecursiveMutex) Lock() {
	gid := goid.Get()

	if atomic.LoadInt64(&rm.owner) == gid {
		rm.recursion++
		return
	}

	rm.Mutex.Lock()

	atomic.StoreInt64(&rm.owner, gid)
	rm.recursion = 1
}

func (rm *RecursiveMutex) Unlock() {
	gid := goid.Get()

	if atomic.LoadInt64(&rm.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d) %d", rm.owner, gid))
	}

	rm.recursion--
	if rm.recursion != 0 {
		return
	}

	atomic.StoreInt64(&rm.owner, -1)

	rm.Mutex.Unlock()
}
