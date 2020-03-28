package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var result int
	if r := maxDepth((*root).Left); result < r {
		result = r
	}
	if r := maxDepth((*root).Right); result < r {
		result = r
	}
	return result + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	r := maxDepth(root.Left)
	l := maxDepth(root.Right)
	if diff := r - l; diff < -1 || 1 < diff {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
