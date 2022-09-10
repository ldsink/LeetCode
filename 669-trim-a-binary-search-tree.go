package main

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var trim func(*TreeNode) *TreeNode
	trim = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val < low || high < node.Val {
			if n := trim(node.Left); n != nil {
				return n
			}
			return trim(node.Right)
		}
		node.Left, node.Right = trim(node.Left), trim(node.Right)
		return node
	}

	return trim(root)
}
