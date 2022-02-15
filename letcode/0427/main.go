// @Description _427
// @Author caopengfei 2021/4/27 14:04
package main

import "fmt"

//处理一个字符数组，只保留一个空格，并返回数组的长度

const _repeat = 1

func main() {
	nums := make([]int, 0)
	for i := 0; i < 10; i++ {
		nums = append(nums, i%3)
	}
	count := handleArr(nums)

	fmt.Println(count)

}

func handleArr(arr []int) int {
	size := len(arr)
	count := len(arr)
	flag := false

	for i := 0; i < size; i++ {
		if arr[i] == _repeat {
			if !flag {
				flag = true
				continue
			}
			for j := i; j < size-1; j++ {
				arr[j] = arr[j+1]
				count--
			}
		}
	}
	return count
}
