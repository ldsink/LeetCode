package main

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	left, right := root.Left, root.Right
	flatten(left)
	flatten(right)
	root.Left = nil
	for root.Right = left; root.Right != nil; root = root.Right {
	}
	root.Right = right
}
