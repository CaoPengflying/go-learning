// @Description p2
// @Author caopengfei 2021/4/12 15:10
package p2

import (
	"fmt"
	"go-learning/concurrency/memory_model"
	"go-learning/concurrency/memory_model/p3"
)

var V1_p2 = memory_model.Trace("init v1 p2", 2)
var V2_p2 = memory_model.Trace("init v2 p2", p3.V2_p3)

func init() {
	fmt.Println("init func in p2")
	V1_p2 = 200
}
