package main

func removeElements(head *ListNode, val int) *ListNode {
	for ; head != nil && head.Val == val; head = head.Next {
	}
	if head == nil {
		return nil
	}

	for current := head; current != nil; {
		if current.Next != nil && current.Next.Val == val {
			current.Next = current.Next.Next
			continue
		}
		current = current.Next
	}

	return head
}
