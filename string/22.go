package string

/*
个人思路：和17题类似，字符串的排列组合，尝试dfs
答案：空间复杂度可以再优化，使用指针或者全局变量

二刷还是没写出来，尝试用联想记忆理解dfs/bfs，dfs/bfs求解就像求二叉树的节点组合，dfs类似先序/中序/后序遍历，根据不同情况可传递不同参数到叶节点。
bfs就像层序遍历，利用队列先存储下一层所有节点。
求排列组合的题可尝试dfs/bfs
*/

func generateParenthesis(n int) []string {
	return parenthesisFunc(n, 0, "")
}

func parenthesisFunc(left, right int, prefix string) []string {
	result := make([]string, 0)
	if left > 0 {
		result = append(result, parenthesisFunc(left-1, right+1, prefix+"(")...)
	}
	if right > 0 {
		result = append(result, parenthesisFunc(left, right-1, prefix+")")...)
	}
	if left == 0 && right == 0 {
		result = append(result, prefix)
	}
	return result
}
