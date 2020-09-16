package pool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool_GetObj(t *testing.T) {
	objPool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		obj, err := objPool.GetObj(time.Second * 1)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", obj)
			println(obj)
			err := objPool.ReleaseObj(obj)
			if err != nil {
				t.Error(err)
			}
		}
	}
}
