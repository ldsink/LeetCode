package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	oddHead := head
	evenHead := head.Next
	var oddTail, evenTail, next *ListNode
	for node, isOdd := head, true; node != nil; node, isOdd = next, !isOdd {
		next = node.Next
		node.Next = nil
		if isOdd {
			if oddTail != nil {
				oddTail.Next = node
			}
			oddTail = node
		} else {
			if evenTail != nil {
				evenTail.Next = node
			}
			evenTail = node
		}
	}
	oddTail.Next = evenHead
	return oddHead
}
