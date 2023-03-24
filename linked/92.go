package linked

/*
思路：没解出来，只记住了reverseList使用了指针翻转。这题霜神思路是头插法，找到prev和cur，然后将cur后面的节点一个个插入到prev后面
答案：递归思路来解决reverse很清晰，很推荐看下92的leetcode最高题解
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	newHead := &ListNode{Next: head}
	prev := newHead
	cnt := 0
	for cnt < left-1 {
		prev = prev.Next
		cnt++
	}
	cur := prev.Next
	for count := 0; count < (right - left); count++ {
		temp := prev.Next
		prev.Next = cur.Next
		cur.Next = cur.Next.Next
		prev.Next.Next = temp
	}
	return newHead.Next
}

////////////////////////////////////////////////////////////////////////

func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var successor *ListNode = nil

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}
