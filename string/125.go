package string

import "strings"

/*
个人思路：首尾两个指针向中间逼近，逼近的过程中判断首尾是否相等
答案：提交答案时，输入""空字符串会panic，我在本地运行没这问题，比较了下霜神的答案，i<=j换成了i<j过了
	 我觉得除非leetcode的&&是长与才会有这种情况，不过把=删去了也算优化吧
*/

func isAlphanumeric(c uint8) bool {
	if ('0' <= c && c <= '9') ||
		('a' <= c && c <= 'z') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlphanumeric(s[i]) {
			i++
		}
		for i < j && !isAlphanumeric(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
