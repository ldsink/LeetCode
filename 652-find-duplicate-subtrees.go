package main

import "fmt"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	cache := make(map[string]int)
	var search func(*TreeNode) string
	var result []*TreeNode
	search = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		left, right := search(node.Left), search(node.Right)
		key := fmt.Sprintf("[%dl%sr%s]", node.Val, left, right)
		if cache[key] == 1 {
			result = append(result, node)
		}
		cache[key]++
		return key
	}
	search(root)
	return result
}
