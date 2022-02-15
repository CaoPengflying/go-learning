// @Description mutex
// @Author caopengfei 2021/2/7 16:53
package mutex

import (
	"fmt"
	"sync"
	"testing"
)

func TestRecursiveMutex_Lock(t *testing.T) {
	count := 0
	wg := sync.WaitGroup{}
	wg.Add(10)
	rm := RecursiveMutex{}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				rm.Lock()
				foo(&rm)
				count++
				rm.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func foo(mutex *RecursiveMutex) {
	mutex.Lock()
	fmt.Println("重入锁定,次数:", mutex.recursion)
	mutex.Unlock()
}
