package linked

/*
思路：没有思路
答案：两种方法：递归/翻转链表，递归比较难想到。
翻转链表的思路是：双指针找到链表中点，从中点开始翻转后半，然后再连接链表的两部分。实现起来细节很多，还是挺难的
*/

// ////////////////////Recursive//////////////////////
func reorderListRecursive(head *ListNode) {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}

	reorderListHelper(head, length)
}

func reorderListHelper(head *ListNode, length int) *ListNode {
	if length == 1 {
		outTail := head.Next
		head.Next = nil
		return outTail
	}
	if length == 2 {
		outTail := head.Next.Next
		head.Next.Next = nil
		return outTail
	}
	tail := reorderListHelper(head.Next, length-2)
	outTail := tail.Next
	subHead := head.Next
	head.Next = tail
	tail.Next = subHead
	return outTail
}

// ////////////////////Reverse//////////////////////
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// reverse
	pre := slow
	cur := slow.Next
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	// contact
	p1, p2 := head, slow.Next
	for p1 != slow {
		slow.Next = p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = slow.Next
	}
}
