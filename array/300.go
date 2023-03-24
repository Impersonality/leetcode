package array

/*
思路：没接出来，看到dp想了想，dp[i]是以i结尾数组的结果集，但没找到dp[i]的推导式。
答案：这题dp思路dp[i]关联的不是dp[i-1]，而是dp[0]~dp[i-1]。
	不过最优解是构造tail数组，其中tails[i] 中存储的是所有长度为 i+1 的上升子序列中末尾最小的值。
	在构造tail的过程中，不断的优化tail各个元素的最小值，虽然最后只用到了tail的长度
*/

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	max := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j] > dp[i] {
				dp[i] = dp[j]
			}
		}
		dp[i]++
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func lengthOfLIS1(nums []int) int {
	tail := make([]int, 0)
	tail = append(tail, nums[0])
	end := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > tail[end] {
			tail = append(tail, nums[i])
			end++
		} else {
			left, right := 0, end
			for left < right {
				mid := (left + right) / 2
				if tail[mid] < nums[i] { // tail[mid] >= num[i]
					left = mid + 1
				} else {
					right = mid
				}
			}
			tail[left] = nums[i]
		}
	}
	return end + 1
}
