// @Description gettersAndSetters
// @Author caopengfei 2022/2/7 17:52

// 总是使用get 和set 方法
// go不推荐使用get 和set方法 比如 time.C

package gettersAndSetters

import "time"

func TestTimer() {
	timer := time.NewTimer(time.Second)
	<-timer.C
}

// 如果使用set和get 则属性名称无需使用get前缀

type Customer struct {
	balance int
}

func (customer *Customer) Balance() int {
	return customer.balance
}

func (customer *Customer) SetBalance(balance int) {
	customer.balance = balance
}
