package array

/*
思路：没有思路，考虑过双指针和dp，但是考虑到结果可能不在数据的边界，所以都排除了。其实结果会肯定在边界。
答案：霜神给了dp和非dp，其实是一样的，区别在有没有构建dp[i]，复杂度一样
*/

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	max := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
