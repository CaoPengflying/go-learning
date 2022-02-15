// @Description init_happen_before
// @Author caopengfei 2021/4/12 15:07
package p3

import (
	"fmt"
	"go-learning/concurrency/memory_model"
)

var V1_p3 = memory_model.Trace("init v1 p3", 3)
var V2_p3 = memory_model.Trace("init v2 p3", 3)

func init() {
	fmt.Println("init func in p3")
	V1_p3 = 300
	V2_p3 = 300
}