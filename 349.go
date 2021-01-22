package leetcode

// sort easy 做的几道都是桶排序，所以这道思路上没啥阻碍，corner case也不复杂
func Intersection(nums1 []int, nums2 []int) []int {
	re := make([]int, 0)
	box := make(map[int]bool)
	for _, v := range nums1 {
		box[v] = true
	}
	for _, v := range nums2 {
		value, ok := box[v]
		if ok && value {
			re = append(re, v)
			box[v] = false
		}
	}
	return re
}
