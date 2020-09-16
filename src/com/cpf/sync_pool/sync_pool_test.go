package sync_pool

import (
	"runtime"
	"sync"
	"testing"
)

// 生命周期 下一次gc
func TestSyncPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			println("create a new obj")
			return 100
		},
	}

	v := pool.Get().(int)
	println(v)

	pool.Put(10)
	// gc 会清除pool中的缓存对象
	runtime.GC()
	v1 := pool.Get().(int)
	println(v1)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			println("create a new obj")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			v := pool.Get().(int)
			println(v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
