package leetcode

/*
思路：没写出来，没读懂题意。map存储思路比较容易想到
答案：正解是双指针，一快一慢，快的每次移动两步，慢的一步，如果快的能追上慢的，说明存在环
*/

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
