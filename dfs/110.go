package dfs

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(depth(root.Left)-depth(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(depth(node.Left), depth(node.Right)) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
