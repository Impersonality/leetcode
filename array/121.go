package leetcode

import "math"

/*
思路：看了tag是dp，竟然都没写出来，对dp掌握的很差。
答案：dp[i] = max(dp[i-1], num[i]-min)，每次循环中，需要算出min，这题可以简化下，不需要dp数组，只需要res
*/

func maxProfit(prices []int) int {
	min, profit := math.MaxInt16, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}
	return profit
}
