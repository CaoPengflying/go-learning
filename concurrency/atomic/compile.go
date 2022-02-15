// @Description atomic
// @Author caopengfei 2021/3/10 16:19
package main

const x int64 = 1 + 1<<33

func main() {
	var i = x
	_ = i
}
