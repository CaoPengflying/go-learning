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
	cxt = context.WithValue(cxt, "key1", "value1")
	cxt = context.WithValue(cxt, "key12", "value12")
	cancel()

	go func() {
		defer func() {
			fmt.Println("goroutine exit")
		}()
		select {
		case <-cxt.Done():
			fmt.Println("ctx done")
			return
		default:
			time.Sleep(time.Second)

		}
	}()

	time.Sleep(time.Second)
	time.Sleep(2 * time.Second)
	//context.WithValue(cxt,nil,nil)
	//context.WithCancel(cxt)
	//context.WithTimeout(cxt,1000)

}
