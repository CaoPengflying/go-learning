// @Description p1
// @Author caopengfei 2021/4/12 15:11
package p1

import (
	"fmt"
	"go-learning/concurrency/memory_model"
	"go-learning/concurrency/memory_model/p2"
)

var V1_p1 = memory_model.Trace("init v1p1", p2.V1_p2)
var V2_p1 = memory_model.Trace("init v2p1", p2.V2_p2)



func init() {
	fmt.Println("init func in p1")
}


