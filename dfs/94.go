package dfs

/*
个人思路：按算法来做题的坏处是不用想这题的算法了，只需考虑实现
答案：二叉树的遍历
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
