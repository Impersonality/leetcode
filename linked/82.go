package linked

/*
思路：双指针，正常情况双指针一起走，特殊情况快指针走，然后处理corner case。测试不过，改了几次，最后结果还可以
答案：双指针，两层循环，内层循环快指针走，外层循环慢指针走
	时间结果不理想，不清楚原因，按理说和我写的应该是差不多的，复杂度都是O(n)
*/

// ////////////////////Iterate//////////////////////
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preHead := &ListNode{Next: head}
	prev, cur := preHead, head.Next
	for cur != nil {
		if prev.Next.Val != cur.Val {
			if prev.Next.Next == cur {
				prev = prev.Next
			} else {
				prev.Next = cur
			}
		}
		cur = cur.Next
	}
	if prev.Next.Next != nil && prev.Next.Val == prev.Next.Next.Val {
		prev.Next = nil
	}
	return preHead.Next
}
