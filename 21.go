package leetcode

/*
指针半天想不明白，先理解到这吧：链表中，一个node用两个指针指向，一个可以迭代下去，
另一个会留在原地，用来return或者其他
*/
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//	head := new(ListNode)
//	headCopy := head
//	for l1 != nil && l2 != nil {
//		if l1.Val <= l2.Val {
//			headCopy.Next = l1
//			l1 = l1.Next
//		}else {
//			headCopy.Next = l2
//			l2 = l2.Next
//		}
//		headCopy = headCopy.Next
//	}
//	if l1 == nil {
//		headCopy.Next = l2
//	}else {
//		headCopy.Next = l1
//	}
//	return head.Next
//}

// recursive
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}
