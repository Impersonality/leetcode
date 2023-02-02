package leetcode

/*
思路：简单题没写出来，递归和迭代都没写出来。
答案：解题思路是遍历list，然后建立反向指针，过程中要使用临时遍历保存next
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var temp *ListNode
	for head != nil {
		next := head.Next
		head.Next = temp
		temp = head
		head = next
	}
	return temp
}
