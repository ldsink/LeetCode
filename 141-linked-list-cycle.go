package main

func hasCycle(head *ListNode) bool {
	for ; head != nil; {
		if head == head.Next {
			return true
		}
		head, head.Next = head.Next, head
	}
	return false
}
