// @Description os
// @Author caopengfei 2021/3/2 11:45
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args)!=0{
		fmt.Println(os.Args[0])// args 第一个片 是文件路径
	}
}
