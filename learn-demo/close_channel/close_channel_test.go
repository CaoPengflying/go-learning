package close_channel

import (
	"sync"
	"testing"
)
func producer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			println(i, " insert into ch")
			ch <- i
		}
		wg.Done()
		close(ch)
	}()
}
// 从channel 中取数据时，有两个返回值，第二个返回值为true则表示获取成功，否则为获取失败
func consumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			if val, ok := <-ch; ok {
				println("consumer", val)
			}
		}
		wg.Done()
	}()
}
// 从channel中只取第一个值时，则默认返回0
func consumer1(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			println("consumer1",<-ch)
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	producer(ch, &wg)
	wg.Add(1)
	consumer(ch, &wg)
	wg.Add(1)
	consumer1(ch, &wg)
	wg.Wait()

}
