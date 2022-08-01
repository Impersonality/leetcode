package string

/*
个人思路：滑动窗口，没有想到解法
答案：我写的O(2)的复杂度，滑动窗口应尽量先滑动右侧，这样是O(1)复杂度
*/

func longSubStr(s string) int {
	start, end := 0, 0
	length := 0
	for start < len(s) {
		exist := make(map[byte]bool)
		for end < len(s) && !exist[s[end]] {
			exist[s[end]] = true
			end++
		}
		if end-start > length {
			length = end - start
		}
		start++
		end = start
	}
	return length
}

func lengthOfLongestSubstring2(s string) int {
	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))
	for left < len(s) {
		if idx, ok := indexes[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		indexes[s[left]] = left
		left++
		res = max(res, left-right)
	}
	return res
}

func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++

		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}
