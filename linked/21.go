package linked

/*
思路：做过该题，也有一点印象，比较两个链表当前的值，小的则加入到新链表。想尝试不构建新链表实现，两个链表打散相互连接，实现不了。看了答案后，链表本身就是指针，构建新链表也就是新建了head，实质还是打散了原链表
答案：霜神用的递归，递归时间和空间复杂度都挺高
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 == nil {
		cur.Next = list2
	}
	if list2 == nil {
		cur.Next = list1
	}
	return head.Next
}

/*
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
*/
