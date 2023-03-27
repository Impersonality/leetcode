package linked

/*
思路：难度不大，需注意特殊情况，即使节点为nil，仍有可能存在进位
*/

// ////////////////////Recursive//////////////////////
func addTwoNumbersRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersHelper(l1, l2, 0)
}

func addTwoNumbersHelper(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	var (
		left  *ListNode
		right *ListNode
	)

	sum := carry
	if l1 != nil {
		sum += l1.Val
		left = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		right = l2.Next
	}
	if sum >= 10 {
		carry = 1
	} else {
		carry = 0
	}
	node := &ListNode{Val: sum % 10}
	node.Next = addTwoNumbersHelper(left, right, carry)
	return node
}

// ////////////////////Iterate//////////////////////
func addTwoNumbersIterate(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur, carry := head, 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		sum = sum % 10
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}
	return head.Next
}
