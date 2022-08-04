package string

/*
个人思路：想尝试滑动窗口，没解出来，因为这题的滑动窗口需要变化一下，应该称作中心扩散
答案：最优解法为马拉车，先添加辅佐字符，这样即使是偶数长度的s，中心也是单个字符。
	 中心扩散前先检查一下，因为回文字符串是对称的，如果可以映射到对称的字符串，中心扩散就可以不从0开始
*/

//////////////////////////中心扩散//////////////////////////

var max = ""

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	max = s[:1]
	for i := 0; i < len(s); i++ {
		middleSpread(s, i, i)
		middleSpread(s, i, i+1)
	}
	return max
}

func middleSpread(s string, left, right int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	if right-left-1 > len(max) {
		max = s[left+1 : right]
	}
}

//////////////////////////马拉车//////////////////////////

func longestPalindrome1(s string) string {
	if len(s) <= 1 {
		return s
	}
	sList := make([]byte, 0)
	sList = append(sList, '#')
	for i := 0; i < len(s); i++ {
		sList = append(sList, s[i])
		sList = append(sList, '#')
	}
	maxRight := 0
	center := 0
	maxLen := 1
	begin := 0
	dp := make([]int, len(sList))
	for i := 0; i < len(sList); i++ {
		if i < maxRight {
			dp[i] = min(dp[2*center-i], maxRight-i)
		}
		left, right := i-(dp[i]+1), i+(dp[i]+1)
		for left >= 0 && right < len(sList) && sList[left] == sList[right] {
			left--
			right++
			dp[i]++
		}
		if i+dp[i] > maxRight {
			maxRight = i + dp[i]
			center = i
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			begin = (i - dp[i]) / 2
		}
	}
	return s[begin : begin+maxLen]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
