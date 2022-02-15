// @Description redis
// @Author caopengfei 2022/2/7 16:12

package redis

import "fmt"

func init() {
	fmt.Println("redis.init")
}

func Store(key, value string) error {
	fmt.Println("redis.Store")
	return nil
}
