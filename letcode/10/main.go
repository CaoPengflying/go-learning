package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(minNumber([]int{3, 30, 34, 5, 9}))
}

func minNumber(nums []int) string {
	nums = insertSort(nums)

	res := ""
	for _, v := range nums {
		res = res + strconv.Itoa(v)
	}
	return res
}
func insertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if compare(nums[j], num){
				nums[j+1] = nums[j]
			}else {
				break
			}
		}
		nums[j+1] = num
	}
	return nums
}

func quickSort() {

}

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if compare(nums[j], nums[j+1]) {
				flag := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = flag
			}
		}
	}
	return nums
}

func selectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if compare(nums[i], nums[j]) {
				flag := nums[i]
				nums[i] = nums[j]
				nums[j] = flag
			}
		}
	}
	return nums
}

func compare(num1, num2 int) bool {
	num1Str := strconv.Itoa(num1)
	num2Str := strconv.Itoa(num2)

	res1, _ := strconv.Atoi(num1Str + num2Str)
	res2, _ := strconv.Atoi(num2Str + num1Str)
	return res1 > res2
}
