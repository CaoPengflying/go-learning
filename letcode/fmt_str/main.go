// @Description fmt_str
// @Author caopengfei 2021/4/20 15:55
package main

import "math"

//给你一个字符串形式的电话号码 number 。number 由数字、空格 ' '、和破折号 '-' 组成。 难度：简单
//请你按下述方式重新格式化电话号码。
//首先，删除 所有的空格和破折号。
//其次，将数组从左到右 每 3 个一组 分块，直到 剩下 4 个或更少数字。剩下的数字将按下述规定再分块：
//2 个数字：单个含 2 个数字的块。
//3 个数字：单个含 3 个数字的块。
//4 个数字：两个分别含 2 个数字的块。
//最后用破折号将这些块连接起来。注意，重新格式化过程中 不应该 生成仅含 1 个数字的块，并且 最多 生成两个含 2 个数字的块。

func main() {
	//str := "123 4-5678"
	//
	//fmtRes := fmtStr(str)
	//
	//fmt.Println(fmtRes)
	//
	//str = "123 4-567"
	//
	//fmtRes = fmtStr(str)
	//
	//fmt.Println(fmtRes)
	//
	//str = "1-23-45 6"
	//
	//fmtRes = fmtStr(str)
	//
	//fmt.Println(fmtRes)
	//[2,4,3]
	//[5,6,4]
	//l1 := ListNode{
	//	Val: 9,
	//}
	//l1.Next = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}
	//
	//l2 := ListNode{
	//	Val:  9,
	//	Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}},
	//}
	//l1 := ListNode{
	//	Val: 2,
	//}
	//l1.Next = &ListNode{Val: 4, Next: &ListNode{Val: 3 }}
	//
	//l2 := ListNode{
	//	Val:  5,
	//	Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}},
	//}
	//addTwoNumbers(&l1, &l2)
	reverse(123)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cursor := head

	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		node := ListNode{}
		num1 := 0
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}

		num2 := 0
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}
		node.Val = (num1 + num2 + carry ) % 10

		if (num1 + num2 + carry) / 10 > 0 {
			carry = 1
		}else {
			carry = 0
		}

		cursor.Next = &node
		cursor = cursor.Next
	}
	return head.Next
}

func reverse(x int) int {
	flag := false
	if x < 0 {
		x = -x
		flag = true
	}

	arr := []int{}

	for x > 0 {
		arr = append(arr,x%10)
		x = x / 10
		continue
	}

	sum := 0
	a := 1
	for i:=len(arr)-1; i>=0; i-- {
		sum = sum + a * arr[i]
		a = a * 10
	}
	if flag {
		sum = -sum
	}
	if sum > math.MaxInt32 || sum < -(math.MaxInt32 + 1){
		return 0
	}
	return sum
}