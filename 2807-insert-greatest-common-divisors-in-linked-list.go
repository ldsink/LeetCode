package main

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	result := head
	for curr, next := head, head.Next; next != nil; curr, next = next, next.Next {
		val := gcd(curr.Val, next.Val)
		curr.Next = &ListNode{Val: val}
		curr.Next.Next = next
	}
	return result
}
