package dfs

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	preorderTraversalFunc(root, &res)
	return res
}

func preorderTraversalFunc(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	preorderTraversalFunc(node.Left, res)
	preorderTraversalFunc(node.Right, res)
}
