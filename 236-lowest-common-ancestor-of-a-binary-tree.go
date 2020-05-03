package main

func findNode(root, node *TreeNode) bool {
	if root == node {
		return true
	}
	return (root.Left != nil && findNode(root.Left, node)) || (root.Right != nil && findNode(root.Right, node))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	} else if root.Left == nil {
		return lowestCommonAncestor(root.Right, p, q)
	} else if root.Right == nil {
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
