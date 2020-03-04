package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k < 2 {
		return head
	}
	var result, lastBlock, endOfBlock, previousNode, next *ListNode
	for current, count := head, 1; current != nil; count, current = count+1, next {
		next = (*current).Next

		if count == 1 {
			endOfBlock = current
			previousNode = current
			(*current).Next = nil
			continue
		}

		(*current).Next = previousNode

		if count < k {
			previousNode = current
			continue
		}

		(*endOfBlock).Next = next // end of current block
		if result == nil {
			result = current
		}
		if lastBlock != nil {
			(*lastBlock).Next = current
		}
		lastBlock = endOfBlock
		previousNode = nil
		count = 0
	}

	// reverse rest nodes 
	if previousNode != nil {
		var headOfRest *ListNode
		current := previousNode
		previousNode = nil
		var next *ListNode
		for ; current != nil; current = next {
			headOfRest = current
			next = (*current).Next
			(*current).Next = previousNode
			previousNode = current
		}
		(*lastBlock).Next = headOfRest
	}

	return result
}
