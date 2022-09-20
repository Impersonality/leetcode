package dfs

/*
个人思路：和112类似，递归到叶节点再判断是否满足条件，传递一个数组记录路径。
答案：golang中函数的参数都是值传递，二维数组如果要传递需要使用指针，但是slice append加入的却是指针，也就是append时需要copy一份对象
*/

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	pathSumF(root, sum, make([]int, 0), &res)
	return res
}

func pathSumF(root *TreeNode, sum int, l []int, res *[][]int) {
	if root == nil {
		return
	}
	l = append(l, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		*res = append(*res, append([]int{}, l...))
		return
	}
	pathSumF(root.Left, sum-root.Val, l, res)
	pathSumF(root.Right, sum-root.Val, l, res)
}
