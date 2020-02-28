package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var current *ListNode
	for ; l1 != nil && l2 != nil; l1, l2 = (*l1).Next, (*l2).Next {
		node := ListNode{
			Val:  (*l1).Val + (*l2).Val + carry,
			Next: current,
		}
		carry = node.Val / 10
		node.Val = node.Val % 10
		current = &node
	}
	for ; l1 != nil; l1 = (*l1).Next {
		node := ListNode{
			Val:  (*l1).Val + carry,
			Next: current,
		}
		carry = node.Val / 10
		node.Val = node.Val % 10
		current = &node
	}
	for ; l2 != nil; l2 = (*l2).Next {
		node := ListNode{
			Val:  (*l2).Val + carry,
			Next: current,
		}
		carry = node.Val / 10
		node.Val = node.Val % 10
		current = &node
	}
	for ; !(carry == 0); {
		node := ListNode{
			Val:  carry,
			Next: current,
		}
		carry = node.Val / 10
		node.Val = node.Val % 10
		current = &node
	}

	// reverse order
	var result *ListNode
	for ; current != nil; current = (*current).Next {
		node := ListNode{Val: (*current).Val, Next: result}
		result = &node
	}
	return result
}

func main() {
	// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	node11 := ListNode{Val: 3}
	node12 := ListNode{Val: 4, Next: &node11}
	node13 := ListNode{Val: 2, Next: &node12}
	node21 := ListNode{Val: 4}
	node22 := ListNode{Val: 6, Next: &node21}
	node23 := ListNode{Val: 5, Next: &node22}

	current := addTwoNumbers(&node13, &node23)
	for ; current != nil; current = (*current).Next {
		fmt.Println(*current)
	}
}
