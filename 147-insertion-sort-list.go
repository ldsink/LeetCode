package main

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var result, node, next, prev, curr *ListNode
	for node = head; node != nil; node = next {
		next = node.Next
		node.Next = nil
		if result == nil {
			result = node
		} else if node.Val <= result.Val {
			node.Next = result
			result = node
		} else {
			for prev, curr = result, result.Next; curr != nil; prev, curr = curr, curr.Next {
				if node.Val <= curr.Val {
					prev.Next = node
					node.Next = curr
					break
				}
			}
			if curr == nil {
				prev.Next = node
			}
		}
	}
	return result
}
