package main

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil {
		return 1
	}

	depth := 1
	for node := root; node.Left != nil; node = node.Left {
		depth++
	}
	fullDepth := depth - 1
	stack := make([]*TreeNode, 1<<uint(depth)-1)
	stack[0] = root
	for i, j := 0, 1<<uint(fullDepth-1)-1; i < j; i++ {
		stack[i<<1+1] = stack[i].Left
		stack[i<<1+2] = stack[i].Right
	}

	// count leaf in tree
	count := 1<<uint(fullDepth) - 1
	for i, j := 1<<uint(fullDepth-1)-1, count; i < j; i++ {
		if stack[i].Left == nil {
			break
		}
		count++
		if stack[i].Right == nil {
			break
		}
		count++
	}
	return count
}
