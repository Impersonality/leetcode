package string

/*
个人思路：没有想到解法，想尝试滑动窗口
答案：这题很显然的dp问题，实现时有个bug困扰我很久：<26如何表示。我想的是十位<=2，个位<=6，这样17-19都无法正确判断，应该使用十位*10+个位<=26
*/

func numDecodings(s string) int {
	if len(s) < 1 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		cur := s[i-1] - '0'
		if cur != 0 {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2]-'0' != 0 && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(dp)-1]
}
