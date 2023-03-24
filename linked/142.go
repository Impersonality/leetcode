package linked

/*
思路：141的进阶，这题如果不限制空间复杂度非常简单，限制O(1)空间想不出来
答案：追及问题求路程，fast和slow第一次相交点是有含义的，利用该点去找到环的起点
*/

/*
fast 的 t = (x1 + x2 + x3 + x2) / 2
slow 的 t = (x1 + x2) / 1

x1 + x2 + x3 + x2 = 2 * (x1 + x2)
所以 x1 = x3
*/

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
