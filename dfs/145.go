package dfs

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	postorderTraversalFunc(root, &res)
	return res
}

func postorderTraversalFunc(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	postorderTraversalFunc(node.Left, res)
	postorderTraversalFunc(node.Right, res)
	*res = append(*res, node.Val)
}
