package leetcode

/*
思路：想到了优先列队的思想，但是go的集合很少，需自己实现。
答案：总是想着用最优解，其实这题答案就是基于merge2List。不同的解法是merge时使用两两merge(O(k2n))，还是归并merge(O(knlogk))
*/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	res := merge2Lists(lists[0], lists[1])
	for i := 2; i < len(lists); i++ {
		res = merge2Lists(res, lists[i])
	}
	return res
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}

/*
func mergeKLists(lists []*ListNode) *ListNode {
	return merge23(lists, 0, len(lists)-1)
}

func merge23(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	return merge2Lists(merge23(lists, l, mid), merge23(lists, mid+1, r))
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		}else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}
*/
