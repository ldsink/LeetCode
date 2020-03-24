package main

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	if (*head).Val >= x {
		for curr, prev := (*head).Next, head; curr != nil; curr = (*curr).Next {
			if (*curr).Val < x {
				(*prev).Next = (*curr).Next
				(*curr).Next = head
				return partition(curr, x)
			}
			prev = curr
		}
	} else {
		for curr, prev, tail := (*head).Next, head, head; curr != nil; {
			next := (*curr).Next
			if (*curr).Val < x {
				if prev == tail {
					tail = curr
					prev = curr
				} else {
					tailNext := (*tail).Next
					(*tail).Next = curr
					(*curr).Next = tailNext
					tail = curr
					(*prev).Next = next
				}
			} else {
				prev = curr
			}
			curr = next
		}
	}
	return head
}
