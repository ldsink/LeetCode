package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	var result []int
	for ; head != nil; head = head.Next {
		result = append(result, head.Val)
	}
	for i, j := 0, len(result)>>1; i < j; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return result
}
