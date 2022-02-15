// @Description single_flight
// @Author caopengfei 2021/4/14 21:14
package single_flight

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 单飞，能够将查询的并发合并成一个请求，只访问一次，其他的并发请求直接使用第一个请求的返回值
var count int32

func TestSingleFlight(t *testing.T) {
	time.AfterFunc(1*time.Second, func() {
		atomic.AddInt32(&count, -count)
	})
	// 3.412698ms
	// 1.002072002s
	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 1000
		sf  = &singleflight.Group{}
	)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			res, _ := singleFlightGetArticle(sf, 1)
			//res, _ := getArticle(1)
			if res != "article:1" {
				panic("err")
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("同时发起 %d 次请求，耗时: %s", n, time.Since(now))
}

func getArticle(id int64) (string, error) {
	atomic.AddInt32(&count, 1)
	time.Sleep(time.Duration(count) * time.Millisecond)

	return fmt.Sprintf("article:%d", id), nil
}

func singleFlightGetArticle(sf *singleflight.Group, id int64) (string, error) {
	v, err, shared := sf.Do(fmt.Sprintf("%d", id), func() (interface{}, error) {
		return getArticle(id)
	})
	fmt.Println("shared", shared)
	return v.(string), err
}

