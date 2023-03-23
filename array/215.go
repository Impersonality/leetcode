package leetcode

/*
思路：无，没思路
答案：通过快排或者桶排实现，要注意下corner case，有点繁琐，不过耐着性子也写出来了
*/

func findKthLargest(nums []int, k int) int {
	i, j := 0, len(nums)-1
	pos := partition(nums, i, j)
	for pos+1 != k {
		pos = partition(nums, i, j)
		if pos+1 < k {
			i = pos + 1
		} else if pos+1 > k {
			j = pos - 1
		}
	}
	return nums[pos]
}

func partition(nums []int, i, j int) int {
	flag := i
	for i < j {
		for i < j && nums[j] <= nums[flag] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		flag = j
		for i < j && nums[i] > nums[flag] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
		flag = i
	}
	return flag
}
