package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	abs := func(x int) int {
		if x >= 0 {
			return x
		}
		return -x
	}
	var f func(int, int, *TreeNode) int
	f = func(max, min int, node *TreeNode) int {
		if node == nil {
			return 0
		}
		result := abs(max - node.Val)
		if r := abs(min - node.Val); result < r {
			result = r
		}
		if max < node.Val {
			max = node.Val
		}
		if min > node.Val {
			min = node.Val
		}
		if r := f(max, min, node.Left); result < r {
			result = r
		}
		if r := f(max, min, node.Right); result < r {
			result = r
		}
		return result
	}
	result := f(root.Val, root.Val, root.Left)
	if r := f(root.Val, root.Val, root.Right); result < r {
		result = r
	}
	return result
}
