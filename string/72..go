package string

/*
个人思路：
没解出来，现在也没状态。dp题的特点是想到转移方程就很简单，难的就是想出方程。
这题dp[i][j]含义是word[:i+1]到word[:j+1]所需要的step
if word1[i-1] == word2[j-1] {
	dp[i][j] = dp[i-1][j-1]
} else {
	dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
}
*/

func minDistance(word1 string, word2 string) int {
	length1, length2 := len(word1)+1, len(word2)+1
	dp := make([][]int, length1)
	for i := range dp {
		dp[i] = make([]int, length2)
	}
	for i := 0; i < length2; i++ {
		dp[0][i] = i
	}
	for i := 0; i < length1; i++ {
		dp[i][0] = i
	}
	for i := 1; i < length1; i++ {
		for j := 1; j < length2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
		}
	}
	return dp[length1-1][length2-1]
}
