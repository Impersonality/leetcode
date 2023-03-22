package leetcode

/*
思路：先算出链表总长度，即可得出待删除节点的顺序，需处理头节点corner case
答案：这道题有一种特别简单的解法。设置 2 个指针，一个指针距离前一个指针 n 个距离。同时移动 2 个指针，2 个指针都移动相同的距离。当一个指针移动到了终点，那么前一个指针就是倒数第 n 个节点了。
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		length++
	}
	if length == n {
		return head.Next
	}

	order := length - n
	cur = head
	for i := 1; i < order; i++ {
		cur = cur.Next
	}
	if cur.Next == nil {
		return head
	}
	cur.Next = cur.Next.Next
	return head
}

//////////////////////TwoPoints//////////////////////
func removeNthFromEndTwoPoints(head *ListNode, n int) *ListNode {
	newHead := &ListNode{Next: head}
	preSlow, slow, fast := newHead, head, head
	for fast != nil {
		if n <= 0 {
			preSlow = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	preSlow.Next = slow.Next
	return newHead.Next
}
