package dfs

/*
将上层的值作为参数传给递归函数，到叶节点累加值
*/

func sumNumbers(root *TreeNode) int {
	var sum int
	sumNumbersFunc(root, 0, &sum)
	return sum
}

func sumNumbersFunc(node *TreeNode, up int, sum *int) {
	if node == nil {
		return
	}
	up = up*10 + node.Val
	if node.Left == nil && node.Right == nil {
		*sum = *sum + up
	}
	sumNumbersFunc(node.Left, up, sum)
	sumNumbersFunc(node.Right, up, sum)
}
