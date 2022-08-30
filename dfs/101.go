package dfs

/*
个人思路：简单的递归，但是没理解题意，symmetric只是对root，root的左右子树对称
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return f(root.Left, root.Right)
}

func f(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && f(left.Left, right.Right) && f(left.Right, right.Left)
}
