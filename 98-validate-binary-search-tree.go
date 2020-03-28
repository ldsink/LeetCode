package main

func isValid(root *TreeNode, lt, gt int, hasLt, hasGt bool) bool {
	if hasLt && (*root).Val >= lt {
		return false
	}
	if hasGt && (*root).Val <= gt {
		return false
	}
	if (*root).Left != nil {
		if (*(*root).Left).Val >= (*root).Val {
			return false
		}
		if !isValid((*root).Left, (*root).Val, gt, true, hasGt) {
			return false
		}
	}
	if (*root).Right != nil {
		if (*(*root).Right).Val <= (*root).Val {
			return false
		}
		if !isValid((*root).Right, lt, (*root).Val, hasLt, true) {
			return false
		}
	}
	return true
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isValid(root, 0, 0, false, false)
}
