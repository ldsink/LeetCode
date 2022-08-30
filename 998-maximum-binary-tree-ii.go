package main

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	var insert func(*TreeNode) *TreeNode
	insert = func(node *TreeNode) *TreeNode {
		if node == nil {
			return &TreeNode{Val: val}
		}
		if val > node.Val {
			return &TreeNode{Val: val, Left: node}
		}
		node.Right = insert(node.Right)
		return node
	}
	return insert(root)
}
