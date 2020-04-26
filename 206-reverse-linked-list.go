package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tail := head
	next := head.Next
	head.Next = nil
	for curr, prev := next, head; curr != nil; curr, prev = next, curr {
		if curr.Next == nil {
			tail = curr
		}
		next, curr.Next = curr.Next, prev
	}
	return tail
}
