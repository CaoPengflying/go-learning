// @Description _422
// @Author caopengfei 2021/4/22 20:40
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nums := []int{1, 4, 5, 6, 7, 3, 2}
	MergeSort(nums, 0, len(nums))
	fmt.Println(nums)
}

func merge(arr []int, l int, mid int, r int) {
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = arr[i]
	}

	left := l
	right := mid + 1

	for i := l; i <= r; i++ {
		if left > mid {
			arr[i] = temp[right-l]
			right++
		} else if right > r {
			arr[i] = temp[left-l]
			left++
		} else if temp[left - l] > temp[right - l] {
			arr[i] = temp[right - l]
			right++
		} else {
			arr[i] = temp[left - l]
			left++
		}
	}
}

func MergeSort(arr []int, l int, r int) {
/*	// 第二步优化，当数据规模足够小的时候，可以使用插入排序
	if r - l <= 15 {
		// 对 l,r 的数据执行插入排序
		for i := l + 1; i <= r; i++ {
			temp := arr[i]
			j := i
			for ; j > 0 && temp < arr[j-1]; j-- {
				arr[j] = arr[j-1]
			}
			arr[j] = temp
		}
		return
	}*/

	mid := (r + l) / 2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)

	// 第一步优化，左右两部分已排好序，只有当左边的最大值大于右边的最小值，才需要对这两部分进行merge操作
	if arr[mid] > arr[mid + 1] {
		merge(arr, l, mid, r)
	}
}



