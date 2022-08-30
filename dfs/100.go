package dfs

/*
个人思路：简单的递归
答案：还可以更简洁一些，比如：return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
*/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) || (p.Val != q.Val) {
		return false
	}
	if !isSameTree(p.Left, q.Left) {
		return false
	}
	if !isSameTree(p.Right, q.Right) {
		return false
	}
	return true
}
