// @Description main
// @Author caopengfei 2021/4/12 15:13
package main

import "sync"

/*func main() {
	fmt.Println("v1p1", p1.V1_p1)
	fmt.Println("v2p1", p1.V2_p1)
}
func init() {
	fmt.Println("init func in main")
}*/

/*var ch = make(chan struct{})
var s string

func f() {
	s = "hello, world"
	<-ch
}

func main() {
	go f()
	ch <- struct{}{}
	print(s)
}*/

// mutex的happens before
/*var mu sync.Mutex
var s string

func f() {
	s = "hello, world"
	mu.Unlock()
}

func main() {
	mu.Lock()
	go f()
	mu.Lock()
	print(s)
}*/

// once 的happens before
var s string
var once sync.Once

func f() {
	s = "hello, world"
}
func main() {
	once.Do(f)
	print(s)
}