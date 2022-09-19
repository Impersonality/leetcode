package dfs

func minDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil {
		return minDepth(node.Right) + 1
	}
	if node.Right == nil {
		return minDepth(node.Left) + 1
	}
	return min(depth(node.Left), depth(node.Right)) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
