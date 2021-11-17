package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

func getSumAndTilt(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftSum, leftTilt := getSumAndTilt(root.Left)
	rightSum, rightTilt := getSumAndTilt(root.Right)
	return leftSum + rightSum + root.Val, abs(leftSum-rightSum) + leftTilt + rightTilt
}

func findTilt(root *TreeNode) int {
	_, result := getSumAndTilt(root)
	return result
}
