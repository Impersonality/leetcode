package leetcode

/*
思路：25没写出来，来写24，看着像纯coding能力题，写的毫无优雅，竟然100%速度ac
答案：霜神的答案非常优雅。思路是两个节点交换转换为3个节点实现，构造一个头节点，头后面的两个节点swap，swap后头结点的next刚好是下一组的头节点
*/

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	var lastTail, temp *ListNode
	for head != nil {
		next := head.Next
		if lastTail != nil {
			if next == nil {
				lastTail.Next = head
				break
			} else {
				lastTail.Next = next
			}
		}
		lastTail = head
		temp = next.Next
		next.Next = head
		head.Next = nil
		head = temp
	}
	return newHead
}

func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for pt := dummy; pt != nil && pt.Next != nil && pt.Next.Next != nil; {
		pt, pt.Next, pt.Next.Next, pt.Next.Next.Next = pt.Next, pt.Next.Next, pt.Next.Next.Next, pt.Next
	}
	return dummy.Next
}
