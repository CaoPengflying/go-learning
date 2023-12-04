// @Description mockMode
// @Author caopengfei 2021/9/7 16:25

package mockMode

type Animal interface {
	Hello(hi string) string
}

type Dog struct {
}

func (d *Dog) Hello(hi string) string {
	return "汪汪 - " + hi
}
