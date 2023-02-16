package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var result *ListNode
	if (*l1).Val > (*l2).Val {
		result = l2
		l2 = (*l2).Next
	} else {
		result = l1
		l1 = (*l1).Next
	}
	current := result
	for l1 != nil && l2 != nil {
		if (*l1).Val > (*l2).Val {
			(*current).Next = l2
			current = l2
			l2 = (*l2).Next
		} else {
			(*current).Next = l1
			current = l1
			l1 = (*l1).Next
		}
	}
	if l1 != nil {
		(*current).Next = l1
	} else if l2 != nil {
		(*current).Next = l2
	}
	return result
}
