// @Description producer_consumer
// @Author caopengfei 2021/4/22 17:44
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
)

func main() {
	ch := make(chan string, 3)

	go func() {
		produce(10, ch)
	}()

	done := make(chan bool)
	go consume(2, ch, done)

	<-done
	close(done)
}

func produce(worker int, ch chan string) {
	var wg sync.WaitGroup
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			count := rand.Intn(10)
			fmt.Println(fmt.Sprintf("producer workder%d prooduce %d msg", i, count))
			for i := 0; i < count; i++ {
				ch <- fmt.Sprintf("%d worker produce msg", i)
			}
		}(i)
	}
	wg.Wait()
	close(ch)
}

func consume(worker int, ch chan string, done chan bool) {
	var wg sync.WaitGroup
	f, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				res, ok := <-ch
				if !ok {
					return
				}
				log.Println(res)
				fmt.Println(res)
			}
		}()
	}
	wg.Wait()
	done <- true
}
