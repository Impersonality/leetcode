package leetcode

import "sort"

/*
思路：明显的双指针，太久没做了。我的实现是3指针，就是多层遍历，只是细节小优化，所以超时了
答案：双指针的意思是先排序，然后两个指针分别从头尾出发，可以减少一层遍历，这题的优化小细节很多
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return ans
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k, target := i+1, len(nums)-1, 0-nums[i]
		for j < k {
			sum := nums[j] + nums[k]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j > i+1 && nums[j] == nums[j-1] {
					j++
				}
				if k < len(nums)-1 && nums[k] == nums[k+1] {
					k--
				}
				j++
				k--
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}
