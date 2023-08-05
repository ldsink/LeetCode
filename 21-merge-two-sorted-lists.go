package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var head, tail *ListNode
	if l1.Val > l2.Val {
		head = l2
		l2 = l2.Next
	} else {
		head = l1
		l1 = l1.Next
	}
	for tail = head; l1 != nil && l2 != nil; {
		if l1.Val > l2.Val {
			tail.Next = l2
			tail = l2
			l2 = l2.Next
		} else {
			tail.Next = l1
			tail = l1
			l1 = l1.Next
		}
	}
	if l1 != nil {
		tail.Next = l1
	} else if l2 != nil {
		tail.Next = l2
	}
	return head
}
