package dfs

/*
个人思路：找到相邻的错误节点，swap。没有理解题意，可能不相邻
答案：bst用中序遍历就是一个有序数组，题目转换为有序数组中存在两个乱序元素，乱序点分为prev和cur，两个元素一定是第一处乱序存在的prev
     和第二处乱序存在的cur
*/

func recoverTree(root *TreeNode) {
	var t1, t2, prev *TreeNode
	_, t1, t2 = inOrderTraverse(t1, t2, root, prev)
	if t1 != nil && t2 != nil {
		t1.Val, t2.Val = t2.Val, t1.Val
	}
}

func inOrderTraverse(t1, t2, node, prev *TreeNode) (*TreeNode, *TreeNode, *TreeNode) {
	if node == nil {
		return prev, t1, t2
	}
	prev, t1, t2 = inOrderTraverse(t1, t2, node.Left, prev)
	if prev != nil && prev.Val > node.Val {
		if t1 == nil {
			t1 = prev
		}
		t2 = node
	}
	prev = node
	prev, t1, t2 = inOrderTraverse(t1, t2, node.Right, prev)
	return prev, t1, t2
}
