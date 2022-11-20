package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestNodes(root *TreeNode, queries []int) [][]int {
	var findMin, findMax func(*TreeNode, int) int
	const failed int = -1
	findMin = func(node *TreeNode, val int) int {
		if node == nil {
			return failed
		}
		if val == node.Val {
			return val
		} else if val < node.Val {
			return findMin(node.Left, val)
		} else if next := findMin(node.Right, val); next != failed {
			return next
		}
		return node.Val
	}
	findMax = func(node *TreeNode, val int) int {
		if node == nil {
			return failed
		}
		if val == node.Val {
			return val
		} else if val > node.Val {
			return findMax(node.Right, val)
		} else if next := findMax(node.Left, val); next != failed {
			return next
		}
		return node.Val
	}
	var result [][]int
	for _, query := range queries {
		result = append(result, []int{findMin(root, query), findMax(root, query)})
	}
	return result
}
