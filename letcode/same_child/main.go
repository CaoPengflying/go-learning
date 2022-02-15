// @Description same_child
// @Author caopengfei 2021/5/5 19:59
package main

func main() {
	a := &TreeNode{Val: 1}
	b := &TreeNode{Val: 2,Left: &TreeNode{Val: 3}}

	isSubStructure(a,b)
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	pa := dfs(A)
	pb := dfs(B)
	if len(pa) < len(pb) {
		return false
	}
	for i := 0; i <= len(pa)-len(pb); i++ {
		flag := true
		for j := 0; j < len(pb); j++ {
			if pa[j] != pb[j] {
				flag = false
				break
			}
		}
		if flag {
			return flag
		}
	}
	return false
}
func dfs(root *TreeNode)(path []int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	leftPath := dfs(root.Left)
	if len(leftPath) > 0 {
		path = append(path,leftPath...)
	}
	rightPath := dfs(root.Right)
	if len(rightPath) > 0 {
		path = append(path,rightPath...)
	}
	return
}
