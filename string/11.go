package string

/*
个人思路：穷举，没有想到解法
答案：指针对撞，求数组的两数之和，三数之和，有序数组的两数运算等等，都可以用指针对撞
*/

func maxArea(height []int) int {
	start, end := 0, len(height)-1
	high := 0
	max := 0
	for start < end {
		width := end - start
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		if max < width*high {
			max = width * high
		}
	}
	return max
}
