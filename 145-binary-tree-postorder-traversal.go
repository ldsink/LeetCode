package main

func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left == nil && node.Right == nil {
			result = append(result, node.Val)
			continue
		}

		stack = append(stack, node)
		if node.Right != nil {
			stack = append(stack, node.Right)
			node.Right = nil
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			node.Left = nil
		}
	}
	return result
}
