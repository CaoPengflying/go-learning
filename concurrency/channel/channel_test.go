// @Description channel
// @Author caopengfei 2021/4/2 16:43
package channel

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestMultiGoroutine(t *testing.T) {
	chs := []chan int{make(chan int, 1), make(chan int, 1), make(chan int, 1)}

	for i := 0; i < 3; i++ {
		go func(ch chan int, i int) {
			time.Sleep(time.Second * time.Duration(i+1))
			fmt.Printf("go%d", i)
			ch <- 1
			close(ch)
		}(chs[i], i)
	}
	select {
	case <-chs[0]:
		fmt.Println("0 end")
	case <-chs[1]:
		fmt.Println("1 end")
	case <-chs[2]:
		fmt.Println("2 end")
	}
	time.Sleep(time.Second * 10)
	fmt.Println("exit")
}

func TestChannel(t *testing.T) {
	ch := make(chan int)
	ch = nil
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("超时")
	}
}

//channel总共有5中用法

// 1. 数据交流 当作并发的 buffer 或者 queue，解决生产者 - 消费者问题。多个 goroutine 可以并发当作生产者（Producer）和消费者（Consumer）。

// 2. 数据传递 一个 goroutine 将数据交给另一个 goroutine，相当于把数据的拥有权 (引 用) 托付出去。

// 3. 信号通知 一个 goroutine 可以将信号 (closing、closed、data ready 等) 传递给另一 个或者另一组 goroutine 。

// 4. 任务编排 可以让一组 goroutine 按照一定的顺序并发或者串行的执行，这就是编排的 功能。

// 5. 锁 利用 Channel 也可以实现互斥锁的机制。

func Test(t *testing.T) {
	intChan := make(chan int, 10)

	go func() {
		for i := range intChan {
			fmt.Println(i)
		}
	}()

	for i := 0; i < 10; i++ {
		intChan <- i
	}

	time.Sleep(10 *time.Second)

}

func TestErr(t *testing.T) {
	ch := make(chan int64, 0)
	// 1. close 为nil 的chan
	ch = nil
	close(ch)
	// 2. send 已经close的chan
	ch <- 1

	// 3。 close 已经close的chan
	close(ch)
}

func TestTaskCompose(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	for {
		go func() {
			ch1 <- 1
		}()
		go func() {
			ch2 <- 2
		}()
		go func() {
			ch3 <- 3
		}()
		go func() {
			ch4 <- 4
		}()

		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
		fmt.Println(<-ch3)
		fmt.Println(<-ch4)
	}
}

// 数据传递
type Token struct {
}

func TestTaskComposeToken(t *testing.T) {
	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}
	for i := 0; i < 4; i++ {
		go newWorker(i, chs[i], chs[(i+1)%4])
	}
	chs[0] <- Token{}
	select {}
}

func newWorker(i int, ch chan Token, nextCh chan Token) {
	for {
		token := <-ch

		fmt.Println(i + 1)
		time.Sleep(time.Second)

		nextCh <- token
	}
}

func TestCases(t *testing.T) {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	cases := createCases(ch1, ch2)

	for i := 0; i < 10; i++ {
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() {
			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
		} else {
			fmt.Println("send:", cases[chosen].Dir, ok)
		}
	}
}

func createCases(chans ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	for _, ch := range chans {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	for i, ch := range chans {
		v := reflect.ValueOf(i)
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: v,
		})
	}
	return cases
}

// 信号通知

func TestNotify(t *testing.T) {
	closing := make(chan struct{})
	closed := make(chan struct{})

	go func() {
		for {
			select {
			case <-closing:
				fmt.Println("业务处理完成")
				return
			default:
				fmt.Println("正在处理业务")
				time.Sleep(time.Second)
			}
		}
	}()

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	close(closing)

	go clearUp(closed)

	select {
	case <-closed:
		fmt.Println("安全释放")
	case <-time.After(100):
		return
	}
	fmt.Println("优雅退出")
}

func clearUp(closed chan struct{}) {
	time.Sleep(20)
	close(closed)
}

//锁
type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := Mutex{
		ch: make(chan struct{}, 1),
	}
	mu.ch <- struct{}{}
	return &mu
}

func (m *Mutex) Lock() {
	<-m.ch
}

func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("error")

	}
}

func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
		return false
	}
}

func TestLock(t *testing.T) {
	mutex := NewMutex()
	flag := 1
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			fmt.Println(flag)
			flag++
			mutex.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(flag)
}

// orDo模式 多个任务，任意一个任务完成则完成
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})

	go func() {
		defer close(orDone)

		var cases []reflect.SelectCase
		for _, ch := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			})
		}

		reflect.Select(cases)
	}()

	return orDone
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)

		time.Sleep(after)
	}()

	return c
}

func TestOrDo(t *testing.T) {
	start := time.Now()

	<-or(
		sig(40*time.Second),
		sig(35*time.Second),
		sig(30*time.Second),
		sig(29*time.Second),
		sig(28*time.Second),
		sig(20*time.Second),
	)

	fmt.Printf("done after %v", time.Since(start))
}

// 扇入模式
func TestFanInReflect(t *testing.T) {
	ch1 := make(chan interface{}, 10)
	ch2 := make(chan interface{}, 10)
	ch3 := make(chan interface{}, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			ch2 <- i * 10
			ch3 <- i * 100
		}
	}()
	out := FanInReflect(ch1, ch2, ch3)

	fmt.Println(<-out)
	select {}
}

func FanInReflect(chans ...chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v
		}
	}()
	return out
}

func TestMaxGoroutineNum(t *testing.T) {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		ch <- 1
		go func() {
			defer func() {
				wg.Done()
				<-ch
			}()
		}()
	}
	wg.Wait()
}
