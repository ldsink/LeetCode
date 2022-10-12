package main

func numComponents(head *ListNode, nums []int) int {
	idx := make(map[int]bool)
	for _, num := range nums {
		idx[num] = true
	}
	var result int
	for inSet := false; head != nil; head = head.Next {
		if !idx[head.Val] {
			inSet = false
			continue
		}
		if !inSet {
			result++
			inSet = true
		}
	}
	return result
}
