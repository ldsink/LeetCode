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
