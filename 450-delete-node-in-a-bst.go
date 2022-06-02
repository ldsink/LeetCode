package main

func deleteNode(root *TreeNode, key int) *TreeNode {
	var arrangeNode func(*TreeNode, *TreeNode)
	arrangeNode = func(node *TreeNode, target *TreeNode) {
		if target.Val < node.Val {
			if node.Left == nil {
				node.Left = target
				return
			}
			arrangeNode(node.Left, target)
		} else {
			if node.Right == nil {
				node.Right = target
				return
			}
			arrangeNode(node.Right, target)
		}
	}

	var deleteTreeNode func(*TreeNode, *TreeNode, bool)
	deleteTreeNode = func(node *TreeNode, parent *TreeNode, isLeft bool) {
		if node == nil {
			return
		} else if node.Val == key {
			if parent == nil { // 删除根节点的情况
				if node.Left == nil {
					root = node.Right
					return
				} else if node.Right == nil {
					root = node.Left
					return
				}
				root = node.Left
				if root.Right == nil {
					root.Right = node.Right
				} else {
					target := root.Right
					root.Right = node.Right
					arrangeNode(root.Right, target)
				}
				return
			} else { // 删除中间节点的情况
				if node.Left == nil {
					if isLeft {
						parent.Left = node.Right
					} else {
						parent.Right = node.Right
					}
					return
				} else if node.Right == nil {
					if isLeft {
						parent.Left = node.Left
					} else {
						parent.Right = node.Left
					}
					return
				}
				if isLeft {
					parent.Left = node.Left
				} else {
					parent.Right = node.Left
				}
				newNode := node.Left
				if newNode.Right == nil {
					newNode.Right = node.Right
				} else {
					target := newNode.Right
					newNode.Right = node.Right
					arrangeNode(newNode.Right, target)
				}
				return
			}
		} else if node.Val < key {
			deleteTreeNode(node.Right, node, false)
		} else {
			deleteTreeNode(node.Left, node, true)
		}
	}

	deleteTreeNode(root, nil, true)
	return root
}
