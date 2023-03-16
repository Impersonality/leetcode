package dfs

/*
答案：完全没思路，看了提示：回溯法，再看代码：完全是dfs穷举。
不过自己再次写也没写出来，要思路非常清晰，才能bug free。
nums切分成两个子数组，一个是走过的路径，一个是未走过的，使用dfs递归实现
*/

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	dfs(nums, []int{}, &res)
	return res
}

func dfs(nums []int, path []int, res *[][]int) {
	c := make([]int, len(path))
	copy(c, path)
	if len(nums) == 0 {
		*res = append(*res, c)
	}
	for i := range nums {
		remainder := make([]int, 0)
		remainder = append(remainder, nums[:i]...)
		remainder = append(remainder, nums[i+1:]...)
		dfs(remainder, append(c, nums[i]), res)
	}
}
