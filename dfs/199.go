package dfs

/*
个人思路：数组的下标就是第i层的result，中根遍历树
答案：霜神用的是bfs，层序遍历，将每层的节点放入一个队列，然后取队列最后一个计算result，然后再舍去上层的节点
*/

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	rightSideViewFunc(root, 0, &res)
	return res
}

func rightSideViewFunc(node *TreeNode, depth int, arr *[]int) {
	if node == nil {
		return
	}
	if len(*arr) <= depth {
		*arr = append(*arr, node.Val)
	}
	rightSideViewFunc(node.Right, depth+1, arr)
	rightSideViewFunc(node.Left, depth+1, arr)
}

// bfs
func rightSideView1(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, queue[n-1].Val)
		queue = queue[n:]
	}
	return res
}
