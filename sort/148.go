package sort

/*
秒看答案233
ListNode排序，147是指定用insertion，链表更适合mergeSort
常规mergeSort使用递归，合并ListNode就是21题，
递归会增加栈空间，空间复杂度log(n)，官答介绍了自底向上的mergeSor，就是遍历，从1开始两两Node merge
看了下实现，理解很困难，先暂时放弃

2刷，忘了，没写出来。抄了遍快排和自底向上的解法
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// //////////////////mergeSort/////////////////////
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

func sortList0(head *ListNode) *ListNode {
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
	return mergeList(sortList0(left), sortList0(midTemp))
}

// //////////////////quickSort/////////////////////
func sortList1(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	quickSort(pre, nil)
	return pre.Next
}

func quickSort(pre, tail *ListNode) {
	if pre == tail || pre.Next == tail || pre.Next.Next == tail {
		return
	}
	pivot := pre.Next
	l := &ListNode{}
	p1 := pre.Next
	p2 := l
	for p1.Next != tail {
		if p1.Next.Val < pivot.Val {
			p2.Next = p1.Next
			p2 = p2.Next
			p1.Next = p1.Next.Next
			p2.Next = nil
		} else {
			p1 = p1.Next
		}
	}
	p2.Next = pre.Next
	pre.Next = l.Next
	quickSort(pre, pivot)
	quickSort(pivot, tail)
}
