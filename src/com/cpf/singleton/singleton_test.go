package singleton

import (
	"sync"
	"testing"
)

type Singleton struct {
}

var singleInstance *Singleton

var once sync.Once

func GetSingleton() *Singleton {

	once.Do(func() {
		println("create singleton")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingleton(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			obj := GetSingleton()
			println(obj)
			wg.Done()
		}()
	}
	wg.Wait()
}
