// @Description context
// @Author caopengfei 2021/3/10 15:04
package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	cxt, cancel := context.WithCancel(context.Background())

	go func() {
		defer func() {
			fmt.Println("goroutine exit")
		}()
	}()

	for {
		select {
		case <-cxt.Done():
			return
		default:
			time.Sleep(time.Second)

		}
	}
	time.Sleep(time.Second)
	cancel()
	time.Sleep(2 * time.Second)
	//context.WithValue(cxt,nil,nil)
	//context.WithCancel(cxt)
	//context.WithTimeout(cxt,1000)

}
