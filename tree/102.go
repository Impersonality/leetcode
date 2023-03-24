package tree

/*
思路：层序遍历思想是用将节点顺序入队列，然后先进先出，将节点的左右子节点入队。我用了go的list，效率很低
答案：霜神的bfs用list作为队列，每次取完一层节点后，截取list的后半部分。也可以用dfs来做这题，递归时传level参数，level可对应到res数组的下标
*/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*TreeNode{root}
	res := make([][]int, 0)

	for len(q) > 0 {
		temp := make([]int, 0)
		qLength := len(q)
		for i := 0; i < qLength; i++ {
			node := q[i]
			temp = append(temp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, temp)
		q = q[qLength:]
	}
	return res
}

// DFS
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	var dfsLevel func(node *TreeNode, level int)
	dfsLevel = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(res) == level {
			res = append(res, []int{node.Val})
		} else {
			res[level] = append(res[level], node.Val)
		}
		dfsLevel(node.Left, level+1)
		dfsLevel(node.Right, level+1)
	}
	dfsLevel(root, 0)
	return res
}
