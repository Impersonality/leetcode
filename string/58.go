package string

/*
个人思路：从后往前找到第一个word的end和start
答案：这题思路很简单，但corner case需要细致处理。easy题我应该耐心些，先写出能完备corner case的解法再优化实现
*/

func lengthOfLastWord(s string) int {
	end := len(s) - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}

	if end < 0 {
		return 0
	}

	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}
	return end - start
}
