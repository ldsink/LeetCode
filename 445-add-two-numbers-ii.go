package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var num1, num2, sum []int
	for ; l1 != nil; l1 = l1.Next {
		num1 = append(num1, l1.Val)
	}
	for i, j := 0, len(num1)-1; i < j; i, j = i+1, j-1 {
		num1[i], num1[j] = num1[j], num1[i]
	}
	for ; l2 != nil; l2 = l2.Next {
		num2 = append(num2, l2.Val)
	}
	for i, j := 0, len(num2)-1; i < j; i, j = i+1, j-1 {
		num2[i], num2[j] = num2[j], num2[i]
	}
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	var i, carry int
	for ; i < len(num2); i++ {
		val := num1[i] + num2[i] + carry
		carry = val / 10
		sum = append(sum, val%10)
	}
	for ; i < len(num1); i++ {
		val := num1[i] + carry
		carry = val / 10
		sum = append(sum, val%10)
	}
	for ; carry != 0; carry /= 10 {
		sum = append(sum, carry%10)
	}
	var head *ListNode
	head = &ListNode{Val: sum[len(sum)-1]}
	for i, prev := len(sum)-2, head; i >= 0; i-- {
		curr := &ListNode{Val: sum[i]}
		prev.Next, prev = curr, curr
	}
	return head
}
