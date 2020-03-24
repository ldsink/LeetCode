package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newHead, tail *ListNode
	var value int
	current := head

	for value := (*head).Val; current != nil; current = (*current).Next {
		if current != head && (*current).Val == value {
			continue
		}
		value = (*current).Val
		if (*current).Next == nil || (*current).Val != (*(*current).Next).Val {
			newHead = current
			tail = current
			break
		}
	}
	if current == nil {
		return nil
	}

	for current = (*current).Next; current != nil; current = (*current).Next {
		if (*current).Val == value {
			continue
		}
		value = (*current).Val
		if (*current).Next == nil || (*current).Val != (*(*current).Next).Val {
			(*tail).Next = current
			tail = current
		}
	}
	(*tail).Next = nil
	return newHead
}
