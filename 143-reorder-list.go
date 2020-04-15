package main

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// find middle
	var fast, slow *ListNode
	for fast, slow = head, head; fast.Next != nil && fast.Next.Next != nil; fast, slow = fast.Next.Next, slow.Next {
	}
	if fast.Next != nil {
		fast = fast.Next
	}

	// reverse
	for prev, node := slow.Next, slow.Next.Next; node != nil; {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	slow.Next.Next = nil
	slow.Next = nil

	// reorder
	for left, right := head, fast; left != nil && right != nil; {
		leftNext, rightNext := left.Next, right.Next
		left.Next, right.Next = right, leftNext
		left, right = leftNext, rightNext
	}
}
