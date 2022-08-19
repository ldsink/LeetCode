package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := TreeNode{}
	idx := 0
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max, idx = nums[i], i
		}
	}
	root.Val = nums[idx]
	if idx > 0 {
		root.Left = constructMaximumBinaryTree(nums[:idx])
	}
	root.Right = constructMaximumBinaryTree(nums[idx+1:])
	return &root
}
