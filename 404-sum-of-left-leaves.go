package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getLeftSum(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	var result int
	if root.Left != nil {
		result += getLeftSum(root.Left)
	}
	if root.Right != nil {
		result += sumOfLeftLeaves(root.Right)
	}
	return result
}

func sumOfLeftLeaves(root *TreeNode) int {
	var result int
	if root.Left != nil {
		result += getLeftSum(root.Left)
	}
	if root.Right != nil {
		result += sumOfLeftLeaves(root.Right)
	}
	return result
}
