package main

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	var prev, tail, reverseHead *ListNode
	for current, i := head, 1; current != nil; i++ {
		next := (*current).Next
		if i+1 == m {
			reverseHead = current
		} else if m <= i && i <= n {
			if i == m {
				tail = current
			}
			if i == n {
				(*tail).Next = next
				if m == 1 {
					head = current
				} else {
					(*reverseHead).Next = current
				}
			}
			(*current).Next = prev
			prev = current
		}
		current = next
	}
	return head
}
