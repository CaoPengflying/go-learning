package main

import (
	"fmt"
	"os"
)

func main() {
	//主程序必须是package main
	//主程序方法名必须为main，并且无返回值，无入参

	//使用os.Args可以作为Go运行的入参
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("hello world!", os.Args[1])
	}
	//使用os.Exit作为Go退出的状态
	os.Exit(-1)
}
