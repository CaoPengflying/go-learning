// @Description merge_sort
// @Author caopengfei 2021/4/24 10:39
package main

import "fmt"

func main() {
	//nums := mergeSort([]int{1, 5, 3, 2, 7, 6})
	//println(nums)

	head := &ListNode{Value: 1, Next: &ListNode{Value: 5, Next: &ListNode{Value: 3, Next: &ListNode{Value: 2}}}}

	newHead := LinkMergeSort(head)
	fmt.Println(newHead)
}


/**
归并排序（MERGE-SORT）是建立在归并操作上的一种有效的排序算法,该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
时间复杂度为：O(n*log(n))
java 中Arrays.sort() ，就是用归并排序来实现的.
*/

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	left := mergeSort(arr[0:i])
	right := mergeSort(arr[i:])
	result := merge(left, right)
	return result
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0 //
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
			continue
		}
		result = append(result, left[m])
		m++
	}
	result = append(result, right[n:]...)
	result = append(result, left[m:]...)
	return result
}

type ListNode struct {
	Value int
	Next  *ListNode
}

func LinkMergeSort(node *ListNode) *ListNode {
	if node.Next == nil {
		return node
	}

	midNode := getLinkMid(node)

	left := LinkMergeSort(node)
	right := LinkMergeSort(midNode)

	return LinkMerge(left, right)
}

func LinkMerge(left *ListNode, right *ListNode) *ListNode {
	head := &ListNode{}
	cursor := head
	for left != nil && right != nil {
		node := ListNode{Value: left.Value}
		if node.Value > right.Value {
			node.Value = right.Value
			right = right.Next
			cursor.Next = &node
			cursor = cursor.Next
			continue
		}
		left = left.Next
		cursor.Next = &node
		cursor = cursor.Next
	}

	if left != nil {
		cursor.Next = left
	}
	if right != nil {
		cursor.Next = right
	}

	return head.Next
}

func getLinkMid(node *ListNode) (right *ListNode) {
	slow := node
	fast := node
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	right = slow.Next
	slow.Next = nil
	return right
}
