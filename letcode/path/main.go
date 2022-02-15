// @Description path
// @Author caopengfei 2021/4/25 21:37
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//[3,5,1,6,2,0,8,null,null,7,4]
func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}}
	lowestCommonAncestor(root, &TreeNode{Val: 5}, &TreeNode{Val: 4})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var getPath func(root, p *TreeNode)
	res := [][]*TreeNode{}
	path := []*TreeNode{}
	getPath = func(root, p *TreeNode) {
		if root == nil {
			return
		}

		path = append(path, root)

		if root.Val == p.Val {
			copyPath := make([]*TreeNode,len(path))
			copy(copyPath, path)
			res = append(res, copyPath)
		}

		getPath(root.Left, p)
		getPath(root.Right, p)
		if len(path) > 0 {
			path = path[:len(path)-1]
		}
	}

	getPath(root, p)
	getPath(root, q)

	i := 0

	for i < len(res[0])-1 && i < len(res[1])-1 {
		if res[0][i+1].Val != res[1][i+1].Val {
			return res[0][i]
		}
		i++
	}
	return res[0][i]
}
