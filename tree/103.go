package tree

/*
思路：102的进阶，层序遍历，reverse间隔层的list。时间30%，空间还可以
答案：dfs（递归）实现，代码简洁点，时间空间都不理想
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	pos := true
	for len(q) > 0 {
		length := len(q)
		temp := make([]int, length)
		for i := 0; i < length; i++ {
			temp[i] = q[i].Val
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		if pos {
			res = append(res, temp)
			pos = false
		} else {
			res = append(res, reverse103(temp))
			pos = true
		}
		q = q[length:]
	}
	return res
}

func reverse103(l []int) []int {
	length := len(l)
	res := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		res[i] = l[length-i-1]
	}
	return res
}
