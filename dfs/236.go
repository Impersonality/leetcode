package dfs

/*
思路：看着和235类似，思路却不同，没想到解法
答案：树的题目还是先考虑递归，边界判断是：node是否==p或者q。这个函数实际判断的是树是否包含p,q，如果左子树包含，并且右子树也包含，则node就是答案
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Left, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}
