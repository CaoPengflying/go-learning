// @Description rank
// @Author caopengfei 2021/9/26 18:22

package rank

import (
	"math/rand"
	"time"
)

func Rand4StuIds(ids []int, id int) []int {
	for index, value := range ids {
		if value == id {
			ids = append(ids[0:index], ids[index+1:]...)
			break
		}
	}

	if len(ids) <= 4 {
		return ids
	}

	res := make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		index := rand.Intn(len(ids) - i)
		res = append(res, ids[index])
		ids = append(ids[0:index], ids[index+1:]...)
	}
	return res
}


