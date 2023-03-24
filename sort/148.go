package sort

/*
秒看答案233
ListNode排序，147是指定用insertion，链表更适合mergeSort
常规mergeSort使用递归，合并ListNode就是21题，
递归会增加栈空间，空间复杂度log(n)，官答介绍了自底向上的mergeSor，就是遍历，从1开始两两Node merge
看了下实现，理解很困难，先暂时放弃
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeList(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Val <= right.Val {
		left.Next = mergeList(left.Next, right)
		return left
	} else {
		right.Next = mergeList(right.Next, left)
		return right
	}
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, mid, right := head, head, head
	for right.Next != nil && right.Next.Next != nil {
		mid = mid.Next
		right = right.Next.Next
	}
	midTemp := mid.Next
	mid.Next = nil
	return mergeList(sortList(left), sortList(midTemp))
}
