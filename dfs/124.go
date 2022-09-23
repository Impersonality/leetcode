package dfs

import "math"

/*
个人思路：node的maxPath = max(left, right, left+root, right+root, root)，然后递归
答案：这题细节不少，不能将node的当前maxPath作为递归值，因为path是需要连续的路径，如果当前node是路径，就必须包含node.val
	 其次如果左右节点都作为path后，不能再和上层的node叠加路径。所以最好有个max全局变量一直比较maxPath
*/

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxPath := math.MinInt32
	getPathSum(root, &maxPath)
	return maxPath
}

func getPathSum(node *TreeNode, maxPath *int) int {
	if node == nil {
		return math.MinInt32
	}
	left := getPathSum(node.Left, maxPath)
	right := getPathSum(node.Right, maxPath)
	curMax := max(max(node.Val+left, node.Val+right), node.Val)
	*maxPath = max(max(curMax, node.Val+left+right), *maxPath)
	return curMax
}
