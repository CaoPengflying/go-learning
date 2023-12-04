// @Description mockMode
// @Author caopengfei 2021/9/7 16:31

package mockMode

type Server interface {
	Create(id int) bool
}

type Repository interface {
	Insert(id int) int
}


