package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var tail *ListNode
	var total int
	for current := head; current != nil; current = (*current).Next {
		total++
		if (*current).Next == nil {
			tail = current
		}
	}
	k = k % total
	if k == 0 {
		return head
	}
	(*tail).Next = head
	tail = head
	for i, j := 0, total-k-1; i < j; i++ {
		tail = (*tail).Next
	}
	head = (*tail).Next
	(*tail).Next = nil
	return head
}
