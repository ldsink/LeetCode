package main

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// divide into two part
	//var fast, slow *ListNode
	slow := head
	for fast := head.Next; fast != nil && fast.Next != nil; fast, slow = fast.Next.Next, slow.Next {
	}
	right := sortList(slow.Next)
	slow.Next = nil
	left := sortList(head)

	// merge sorted list
	if left.Val > right.Val {
		head = right
		right = right.Next
	} else {
		head = left
		left = left.Next
	}
	tail := head
	for ; left != nil && right != nil; {
		if left.Val > right.Val {
			tail.Next = right
			right = right.Next
		} else {
			tail.Next = left
			left = left.Next
		}
		tail = tail.Next
	}
	if left != nil {
		tail.Next = left
	} else {
		tail.Next = right
	}
	return head
}
