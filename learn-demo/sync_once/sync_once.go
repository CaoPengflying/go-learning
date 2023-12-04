package sync_once

import (
	"fmt"
	"sync"
)

var once = sync.Once{}

func Fun() {
	once.Do(func() {
		fmt.Println("hello world")
	})
}
