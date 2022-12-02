package dfs

import "fmt"

/*
思路：递归题，边界判断是若root==nil或者root.Left==nil&&root.Right==nil，返回root.val，然后其他节点拼接左子树和右子树的递归结果
答案：
*/

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%d", root.Val)}
	}
	res := make([]string, 0)
	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)
	for _, item := range left {
		res = append(res, fmt.Sprintf("%d->", root.Val)+item)
	}
	for _, item := range right {
		res = append(res, fmt.Sprintf("%d->", root.Val)+item)
	}
	return res
}
