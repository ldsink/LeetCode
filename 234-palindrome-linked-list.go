package main

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	// find middle
	var fast, slow *ListNode
	for fast, slow = head, head; fast.Next != nil && fast.Next.Next != nil; fast, slow = fast.Next.Next, slow.Next {
	}
	if fast.Next != nil {
		fast = fast.Next
	}

	// reverse right part
	var prev, next *ListNode
	for node := slow.Next; node != nil; node = next {
		next = node.Next
		node.Next = prev
		prev = node
	}
	slow.Next = nil

	for ; head != nil && fast != nil; head, fast = head.Next, fast.Next {
		if head.Val != fast.Val {
			return false
		}
	}
	return true
}
