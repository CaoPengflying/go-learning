// @Description main
// @Author caopengfei 2022/2/7 16:12

// main depend on redis
// redis init 最先被执行

package main

import (
	"fmt"
	"go-learning/mistakes100/misusingInit/redis"
	"log"
)

func init() {
	fmt.Println("main.init")
}

func main() {
	err := redis.Store("key", "value")
	fmt.Println("main")
	log.Println(err)
}
