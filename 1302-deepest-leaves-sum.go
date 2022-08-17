package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth, sum int
	var search func(*TreeNode, int)
	search = func(node *TreeNode, d int) {
		if depth < d {
			depth = d
			sum = node.Val
		} else if depth == d {
			sum += node.Val
		}
		if node.Left != nil {
			search(node.Left, d+1)
		}
		if node.Right != nil {
			search(node.Right, d+1)
		}
	}
	search(root, 0)
	return sum
}
