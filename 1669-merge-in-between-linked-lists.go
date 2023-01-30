package main

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	list2Tail := list2
	for ; list2Tail.Next != nil; list2Tail = list2Tail.Next {
	}
	var head, tail *ListNode
	for i, node := 0, list1; node != nil; i, node = i+1, node.Next {
		if i == a-1 {
			head = node
		} else if i == b {
			tail = node.Next
			break
		}
	}
	head.Next = list2
	list2Tail.Next = tail
	return list1
}
