package groutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i <= 1000000; i++ {
		//go func(i int) {
		//	println(i)
		//}(i)
		//tmp := i
		go func() {
			println(i)
		}()
	}
	time.Sleep(3 * time.Second)

}

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestWait(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	//time.Sleep(1 * time.Second)
	wg.Wait()
	t.Logf("counter = %d", counter)

}

func TestA(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			t.Log(i)
		}(i)
	}
	wg.Wait()
}

func TestB(t *testing.T) {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i2: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestC(t *testing.T) {
	m := make(map[int]int)

	c := m[1]
	t.Log(c)
}

func TestD(t *testing.T) {
	t.Log(time.Now().Unix())
}

func TestPrintIndex(t *testing.T) {
	for i := 0; i < 16; i++ {
		fmt.Printf("create index idx_interaction_id on xes_common_record_speak_log_%d (interaction_id);\n", i)
	}
}
