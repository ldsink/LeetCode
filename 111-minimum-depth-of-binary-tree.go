package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if (*root).Left == nil && (*root).Right == nil {
		return 1
	} else if (*root).Left == nil {
		return 1 + minDepth((*root).Right)
	} else if (*root).Right == nil {
		return 1 + minDepth((*root).Left)
	}
	result := minDepth((*root).Left)
	if r := minDepth((*root).Right); result > r {
		result = r
	}
	return result + 1
}
