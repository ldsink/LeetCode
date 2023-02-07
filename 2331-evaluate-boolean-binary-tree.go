package main

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil {
		return root.Val == 1
	}
	left, right := evaluateTree(root.Left), evaluateTree(root.Right)
	if root.Val == 2 {
		return left || right
	}
	return left && right
}
