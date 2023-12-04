package contextConcurrent

import (
	"context"
	"encoding/json"
	"testing"
	"time"
)

func TestReq(t *testing.T) {
	req := mapReq{
		TotalNum:         10,
		StuStatisticInfo: map[int]int{1: 1},
	}
	marshal, err := json.Marshal(req)
	if err != nil {
		return
	}
	t.Log(string(marshal))
}

func TestFun(t *testing.T) {

	go func() {
		time.Sleep(1000)
		return
	}()

	time.Sleep(2000)

	t.Log("end")
}

func TestFun2(t *testing.T) {
	t.Log(time.Now().Unix() - 3600)
	t.Log(time.Now().Unix() + 2700)
	t.Log(time.Now().Unix() > 1680168900)
}

func TestCopyContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "a", "b")
	ctx = context.WithValue(ctx, "x_rpcid", "b")
	reqMap := map[string]string{"a": "b"}
	ctx = context.WithValue(ctx, "hostname", reqMap)
	for i := 0; i < 100; i++ {
		copyContext := CopyContext(ctx)
		go func() {
			value := copyContext.Value("hostname")
			t.Log("1", value)
			//判断value的类型是否为map
			if value != nil {
				value.(map[string]string)["c"] = "d"
			}
			t.Log("2", value)

		}()
	}
	time.Sleep(5 * time.Second)
}
