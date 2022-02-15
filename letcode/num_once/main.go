// @Description num_once
// @Author caopengfei 2021/5/6 17:04
package main

func main() {
	singleNumbers([]int{1, 2, 5, 2})
}

func singleNumbers(nums []int) []int {
	flag := 0

	for _, num := range nums {
		flag ^= num
	}

	div := 1

	for div&flag == 0 {
		div <<= 1
	}

	a := 0
	b := 0
	for _, num := range nums {
		if div&num == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
