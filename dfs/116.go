package dfs

/*
个人思路：先连接root的left和right，再连接root.left.right和root.right.left
答案：只连接两层是不行的，这样3层4层...就没有连接，需要递归连接
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectFunc(root.Left, root.Right)
	return root
}

func connectFunc(left, right *Node) {
	if left == nil {
		return
	}
	left.Next = right
	connectFunc(left.Left, left.Right)
	connectFunc(right.Left, right.Right)
	connectFunc(left.Right, right.Left)
}
