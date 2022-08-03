package string

/*
个人思路：和17题类似，字符串的排列组合，尝试dfs
答案：空间复杂度可以再优化，使用指针或者全局变量
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
