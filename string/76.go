package string

/*
个人思路：没有想到解法，想尝试滑动窗口，但是没想到如果判断子串是否满足条件
答案：使用map或者byte数组 + count 记录子串是否满足条件
*/

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	resultL, resultR := 0, 0
	left, right := 0, 0
	sCnt, tCnt := make(map[byte]int), make(map[byte]int)
	equalCnt := 0
	windowLen := len(s) + 1
	for i := 0; i < len(t); i++ {
		tCnt[t[i]]++
	}
	for left < len(s) {
		if equalCnt < len(t) && right < len(s) {
			sCnt[s[right]]++
			if sCnt[s[right]] <= tCnt[s[right]] {
				equalCnt++
			}
			right++
		} else {
			if equalCnt == len(t) && (right-left < windowLen) {
				windowLen = right - left + 1
				resultL, resultR = left, right
			}
			if sCnt[s[left]] == tCnt[s[left]] {
				equalCnt--
			}
			sCnt[s[left]]--
			left++
		}
	}
	if resultR != 0 {
		return s[resultL:resultR]
	}
	return ""
}
