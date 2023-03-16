package leetcode

/*
思路：简单的解法还是很容易想到，构建map存储一个list的全部节点，list b中若有相同，则为交点
答案：o(m+n) o(1)的解法没想到。如果两个list长度相同，那么很简单，双指针一起走，所以就需要给双链表补齐长度，而且是在链表前端补长度
     list a补上list b, list b 补上list a。如果双链表没有交点，那么会在最后的nil处相等（也可以理解为一个交点）
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for node := headA; node != nil; node = node.Next {
		m[node] = struct{}{}
	}
	for node := headB; node != nil; node = node.Next {
		if _, ok := m[node]; ok {
			return node
		}
	}
	return nil
}

//////////////////////////////////////////////

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
