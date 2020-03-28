package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := [][2]*TreeNode{[2]*TreeNode{(*root).Left, (*root).Right}}
	for ; len(stack) > 0; {
		left, right := stack[0][0], stack[0][1]
		stack = stack[1:]
		if left == nil && right == nil {
			continue
		} else if left == nil && right != nil {
			return false
		} else if left != nil && right == nil {
			return false
		} else if (*left).Val != (*right).Val {
			return false
		}
		stack = append(stack, [2]*TreeNode{(*left).Left, (*right).Right})
		stack = append(stack, [2]*TreeNode{(*left).Right, (*right).Left})
	}
	return true
}
