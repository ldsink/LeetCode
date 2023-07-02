package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var head, prev, curr *ListNode
	for ; l1 != nil && l2 != nil; l1, l2, prev = l1.Next, l2.Next, curr {
		val := l1.Val + l2.Val + carry
		carry = val / 10
		curr = &ListNode{Val: val % 10}
		if head == nil {
			head = curr
		}
		if prev != nil {
			prev.Next = curr
		}
	}
	rest := l1
	if rest == nil {
		rest = l2
	}
	for ; rest != nil; rest, prev = rest.Next, curr {
		val := rest.Val + carry
		carry = val / 10
		curr = &ListNode{Val: val % 10}
		if head == nil {
			head = curr
		}
		if prev != nil {
			prev.Next = curr
		}
	}
	for ; carry != 0; prev = curr {
		curr = &ListNode{Val: carry % 10}
		carry /= 10
		if head == nil {
			head = curr
		}
		if prev != nil {
			prev.Next = curr
		}
	}
	return head
}
