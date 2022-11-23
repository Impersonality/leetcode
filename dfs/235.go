package dfs

/*
思路：如果p<node.val && node.val<q.val，那么val就是最近父辈节点，否则递归左子树或者右子树
答案：
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	if root.Val > p.Val && root.Val < q.Val {
		return root
	}
	if root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return lowestCommonAncestor(root.Right, p, q)
}
