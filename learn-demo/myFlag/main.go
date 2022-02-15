// @Description myFlag
// @Author caopengfei 2021/12/3 15:20

package main

import (
	"flag"
	"fmt"
)


func main() {

	flag.Parse()
	var usr1 = flag.String("usr1", "", "user defined flag -usr1")

	fmt.Println(*usr1)
}
