package leetcode

// 这题和349没啥区别，但是没拿到100是有点惊讶的，去看霜神答案，思路一样
// 细节也接近，然后leetcode测试，效率也差不多，就是跑几次运气好会有100(测试也是不稳定的呀。。。
// 看官方答案，可以用双指针，在两两比较和有序数组时，双指针点赞
func intersect(nums1 []int, nums2 []int) []int {
	re := make([]int, 0)
	box := make(map[int]int)

	for _, v := range nums1 {
		box[v]++
	}

	for _, v := range nums2 {
		value, ok := box[v]
		if ok && value > 0 {
			re = append(re, v)
			box[v]--
		}
	}
	return re
}
