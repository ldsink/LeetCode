package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func constructListNode(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}
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

func printListNode(node *ListNode) {
	for ; node != nil; node = (*node).Next {
		fmt.Print((*node).Val, " -> ")
	}
	fmt.Println("End")
}

func main() {
	fmt.Println("- - - -")
}
