package linked

/*
思路：花了很长时间，并且debug修改了几次才通过。这题思路和24不像，和206 reverseList很像。
	24因为是相邻节点swap，所以将list转换成group list（一组3元素），然后遍历group，group内swap。本题我的思路是，找到k个节点后，然后reverse，再连接头尾.
答案：霜神答案是递归，找到k个节点reverse，然后递归。大体思路是和reverse list差不多
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead, cnt := &ListNode{}, 0

	h, l := head, newHead // h:head  t:tail  l:lastTail
	var t *ListNode

	for head != nil {
		next := head.Next
		if (cnt+1)%k == 0 {
			t = head
			reverseList25(h, t.Next, l)
			h, l, h.Next = next, h, next
		}
		head = next
		cnt++
	}
	return newHead.Next
}

func reverseList25(head, tail, lastTail *ListNode) {
	var temp *ListNode
	for head != tail {
		next := head.Next
		head.Next = temp
		temp = head
		head = next
	}
	lastTail.Next = temp
}

//////////////////////////////////////////////////

func reverseKGroup1(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse25(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse25(first *ListNode, last *ListNode) *ListNode {
	prev := last
	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}
	return prev
}
