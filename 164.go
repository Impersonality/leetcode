package leetcode

/*
个人思路：先排序，然后遍历有序数组找到max diff，时间复杂度nlog(n)+n
		或者insertion，每次插入时记录与前后元素的diff
答案：确实是排序，快排已经100了，还可以用radixSort
*/

func maximumGap1(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	// m is the maximal number in nums
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		m = Max(m, nums[i])
	}
	exp := 1 // 1, 10, 100, 1000 ...
	R := 10  // 10 digits
	aux := make([]int, len(nums))
	for (m / exp) > 0 { // Go through all digits from LSB to MSB
		count := make([]int, R)
		for i := 0; i < len(nums); i++ {
			count[(nums[i]/exp)%10]++
		}
		for i := 1; i < len(count); i++ {
			count[i] += count[i-1]
		}
		for i := len(nums) - 1; i >= 0; i-- {
			tmp := count[(nums[i]/exp)%10]
			tmp--
			aux[tmp] = nums[i]
			count[(nums[i]/exp)%10] = tmp
		}
		for i := 0; i < len(nums); i++ {
			nums[i] = aux[i]
		}
		exp *= 10
	}
	maxValue := 0
	for i := 1; i < len(aux); i++ {
		maxValue = Max(maxValue, aux[i]-aux[i-1])
	}
	return maxValue
}
