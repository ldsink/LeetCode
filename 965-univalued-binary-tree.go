package main

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	val := root.Val
	var isMatch func(*TreeNode) bool
	isMatch = func(node *TreeNode) bool {
		return node.Val == val && (node.Left == nil || isMatch(node.Left)) && (node.Right == nil || isMatch(node.Right))
	}
	return isMatch(root)
}
