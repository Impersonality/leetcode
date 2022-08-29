package dfs

import "math"

func isValidBST(root *TreeNode) bool {
	return support(root, math.Inf(-1), math.Inf(1))
}

func support(node *TreeNode, min, max float64) bool {
	if node == nil {
		return true
	}
	v := float64(node.Val)
	return v > min && v < max && support(node.Left, min, v) && support(node.Right, v, max)
}
