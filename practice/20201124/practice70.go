package main


import (
	"fmt"
	"strconv"
	"sync"
	"time"
)
/**
Snapshot表示复制镜像的意思，但是全部都指向的是同一个map。并没有复制出来镜像。
新的值进行改变时，镜像中的数据也在改变
 */
type Stats struct {
	mutex sync.Mutex

	counters map[string]int
}

func (s *Stats) Snapshot() map[string]int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.counters
}

func (s *Stats) Add(name string, num int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.counters[name] = num
}

func main() {
	stats := Stats{
		counters: map[string]int{},
	}
	for i := 0; i < 100; i++ {
		i := i
		go func() {
			stats.Add(strconv.Itoa(i), i)
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			counters := stats.Snapshot()
			fmt.Println(counters)
		}()
	}
	time.Sleep(10*time.Second)
}
