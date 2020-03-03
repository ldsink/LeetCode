package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func constructListNode(array []int) *ListNode {
	current := ListNode{Val: array[0]}
	result := &current
	currentPtr := result
	for idx, val := range array {
		if idx == 0 {
			continue
		}
		node := ListNode{Val: val}
		(*currentPtr).Next = &node
		currentPtr = &node
	}
	return result
}
