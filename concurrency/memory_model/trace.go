// @Description main
// @Author caopengfei 2021/4/12 15:07
package memory_model

import "fmt"

func Trace(t string, v int) int {
	fmt.Println(t, ":", v)
	return v
}
