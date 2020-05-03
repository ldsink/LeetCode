package main

func findNode(root, node *TreeNode) bool {
	if root == node {
		return true
	}
	return (root.Left != nil && findNode(root.Left, node)) || (root.Right != nil && findNode(root.Right, node))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	if root.Left == nil || root.Right == nil {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			return lowestCommonAncestor(root.Right, p, q)
		}
		return lowestCommonAncestor(root.Left, p, q)
	}

	if findNode(root.Left, p) {
		if findNode(root.Right, q) {
			return root
		}
		return lowestCommonAncestor(root.Left, p, q)
	}

	if findNode(root.Left, q) {
		return root
	}
	return lowestCommonAncestor(root.Right, p, q)
}
