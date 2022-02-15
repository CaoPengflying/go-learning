// @Description _3
// @Author caopengfei 2021/4/24 10:16
package main

var (
	a = A
)

func main() {
	println(search([]int{7, 0, 1, 2, 3, 4, 5, 6}, 7))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2

	for left <= right {
		if target == nums[mid] {
			return mid
		}
		if nums[0] > nums[mid] {
			//右边有序
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1

			} else {
				right = mid - 1
			}
		} else {
			// 左边有序
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		mid = (left + right) / 2
	}
	return -1
}
