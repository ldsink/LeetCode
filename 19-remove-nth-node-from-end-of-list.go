package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var nodePtr []*ListNode
	for ; head != nil; head = (*head).Next {
		nodePtr = append(nodePtr, head)
	}
	// for special case
	if len(nodePtr) == 0 {
		return nil
	}
	if n == 0 {
		return nodePtr[0]
	}
	// remove head
	if n == len(nodePtr) {
		return (*nodePtr[0]).Next
	}
	// remove node
	(*nodePtr[len(nodePtr)-1-n]).Next = (*nodePtr[len(nodePtr)-n]).Next
	return nodePtr[0]
}
