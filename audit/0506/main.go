// @Description _506
// @Author caopengfei 2021/5/7 10:11
package main

func main() {
	s := "123"
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			res += string(s[i])
		}
	}
}

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

//二叉树  跟节点 8 7 6 5 4 3 2

func LevelTraversal(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	q := Queue{}
	q.InQueue(root)
	for !q.IsEmpty() {
		tmp := make([]int, 0)
		for i := q.Count(); i > 0; i-- {
			node := q.DeQueue()
			tmp = append(tmp, node.Value)
			if node.Left != nil {
				q.InQueue(node.Left)
			}
			if node.Right != nil {
				q.InQueue(node.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}

type Queue struct {
	list  []*Node
	count int
}

func (q *Queue) InQueue(node *Node) {
	q.list = append(q.list, node)
	q.count++
}

func (q *Queue) DeQueue() *Node {
	node := q.list[0]
	q.list = q.list[1:]
	q.count--
	return node
}

func (q *Queue) IsEmpty() bool {
	return q.count == 0
}

func (q *Queue) Count() int {
	return q.count
}
