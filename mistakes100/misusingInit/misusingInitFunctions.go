// @Description misusingInit
// @Author caopengfei 2022/2/7 15:50

// 误用 init 函数
// 合理使用init函数
// init函数的缺点
//限制错误管理
//它可能会使如何实现测试变得复杂（例如，要设置外部依赖项，这对于单元测试的范围可能不是必需的）
//如果初始化需要设置状态，则必须通过全局变量来完成

//一般读取静态配置可以使用Init函数

package main

import (
	"fmt"

	// 代码不依赖redis，但是依赖redis中的init，可以通过 _ 引入redis 可以先执行redis中的init
	_ "go-learning/mistakes100/misusingInit/redis"
)

// 当初始化一个包时，所有的常量和变量最先被定义，其次是init
var a = func() int {
	fmt.Println("var")
	return 0
}()

// 多个init函数的调用顺序以源码的先后顺序为准
func init() {
	fmt.Println("init")
}
func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	fmt.Println("main")
	//init function 不能在main中直接调用
	//init()
}
