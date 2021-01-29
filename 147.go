package leetcode

/*
思路很简单，但是一写，超时…
看了答案，加个头指针，然后赋值转换时注意：指针类型是引用类型，copy时是deep copy，会一起发生值变换
对着答案一个个改，才通过，对时间要求好严格。。。
*/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newHead := new(ListNode)
	newHead.Next = head

	cur := head
	next := head.Next

	for next != nil {
		if cur.Val <= next.Val {
			cur = cur.Next
		} else {
			prev := newHead
			for prev.Next.Val <= next.Val {
				prev = prev.Next
			}
			cur.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		next = cur.Next
	}
	return newHead.Next
}
