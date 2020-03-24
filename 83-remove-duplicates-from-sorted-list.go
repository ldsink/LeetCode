package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tail := head
	for current := (*head).Next; current != nil; current = (*current).Next {
		if (*current).Val == (*tail).Val {
			continue
		}
		(*tail).Next = current
		tail = current
	}
	(*tail).Next = nil
	return head
}
