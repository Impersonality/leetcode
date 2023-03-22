package dfs

/*
个人思路：按算法来做题的坏处是不用想这题的算法了，只需考虑实现
答案：二叉树的遍历
	没想到二叉树遍历的迭代方法不简单，我理解是使用栈模拟递归，需要一个标志来表示递归的停止条件
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	search(root, &res)
	return res
}

func search(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	search(node.Left, res)
	*res = append(*res, node.Val)
	search(node.Right, res)
}

//////////////////////Iteration//////////////////////
func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	const (
		White = iota
		Gray
	)

	type Elem struct {
		Color int
		Node  *TreeNode
	}
	res := make([]int, 0)
	q := []*Elem{{Color: White, Node: root}}
	for len(q) > 0 {
		node := q[len(q)-1]
		q = q[:len(q)-1]
		switch node.Color {
		case White:
			if node.Node.Right != nil {
				q = append(q, &Elem{Color: White, Node: node.Node.Right})
			}
			q = append(q, &Elem{Color: Gray, Node: node.Node})
			if node.Node.Left != nil {
				q = append(q, &Elem{Color: White, Node: node.Node.Left})
			}
		case Gray:
			res = append(res, node.Node.Val)
		}
	}
	return res
}
