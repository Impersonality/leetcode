package leetcode

/*
个人思路：遍历数组，先算出前面两个字符串的prefix，然后用prefix与数组中剩下的元素一个个计算prefix
答案：我的思路算是横向遍历，还有纵向遍历：比较所有字符串的第一个char是否相等，依次往后比较。时间复杂度都是O(mn)
	 我的代码还可以优化下，将第一个str设为prefix，然后遍历时不满足条件则修改prefix
*/

func sharedPrefix(a, b string) string {
	lenA, lenB, length := len(a), len(b), 0
	if lenA < lenB {
		length = lenA
	} else {
		length = lenB
	}
	prefix := ""
	for i := 0; i < length; i++ {
		if a[i] == b[i] {
			prefix += string(a[i])
		} else {
			break
		}
	}
	return prefix
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	prefix := sharedPrefix(strs[0], strs[1])
	for i := 2; i < len(strs); i++ {
		prefix = sharedPrefix(prefix, strs[i])
	}
	return prefix
}
