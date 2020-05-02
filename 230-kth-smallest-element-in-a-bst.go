package main

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{root}
	for count := 0; len(stack) > 0; {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left == nil && node.Right == nil {
			count++
			if count == k {
				return node.Val
			}
			continue
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			node.Right = nil
		}
		stack = append(stack, node)
		if node.Left != nil {
			stack = append(stack, node.Left)
			node.Left = nil
		}
	}
	return 0
}
