package leetcode

/*
思路：递归思路很简单，但是更优的解法没想到
答案：先查出树的高度，然后二分查找最后一层。完全二叉树节点数的数转换成二进制，比如6->110，0往左走，1往右走
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
