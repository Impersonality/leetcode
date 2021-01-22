package leetcode

// 思路：遍历一遍，奇偶不对的位置用两个数组记录下来，然后swap这些数
// 答案：双指针，而且细节非常棒，一个指针先动，遇到异常再移动另一指针，for 初始化两个数，golang也可以优雅的swap
func sortArrayByParityII(A []int) []int {
	for i, j := 0, 1; i < len(A); i += 2 {
		if A[i]%2 == 1 {
			for A[j]%2 == 1 {
				j += 2
			}
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}
