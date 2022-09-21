package dfs

/*
个人思路：没有思路
答案：既然是dfs，那么先看最小的情况，先把root.left接到root.right，然后在其后面接上原来的root.right，清空root.left。递归下去
*/

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	temp := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = temp
}
