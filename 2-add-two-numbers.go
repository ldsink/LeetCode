package main

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
