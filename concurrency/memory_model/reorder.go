// @Description memory_model
// @Author caopengfei 2021/4/12 14:37
package memory_model

var a1 string
var done bool

func setup() {
	a1 = "hello, world"
	done = true
}

func main() {
	setup()
	for !done {

	}
	print(a1)
}
