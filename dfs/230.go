package dfs

/*
思路：inorder遍历树，遍历一个node，k-1，如果k==0，那么该node就是要找的node
答案：差不多，设置变量cnt=0，遍历一个node，cnt+=1，若cnt==k，返回val
*/

func kthSmallest(root *TreeNode, k int) int {
	val, _ := func230(root, k)
	return val
}

func func230(node *TreeNode, k int) (int, int) {
	if node.Left != nil {
		val, rest := func230(node.Left, k)
		if rest == 0 {
			return val, rest
		}
		k = rest
	}
	if k == 1 {
		return node.Val, 0
	}
	if node.Right != nil {
		return func230(node.Right, k-1)
	}
	return node.Val, k - 1
}
