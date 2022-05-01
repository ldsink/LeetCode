package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getSortedArray(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := getSortedArray(root.Left)
	result = append(result, root.Val)
	return append(result, getSortedArray(root.Right)...)
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	a, b := getSortedArray(root1), getSortedArray(root2)
	result := make([]int, len(a)+len(b))
	var i, j, k int
	for ; i < len(a) && j < len(b); k++ {
		if a[i] < b[j] {
			result[k] = a[i]
			i++
		} else {
			result[k] = b[j]
			j++
		}
	}
	for ; i < len(a); k++ {
		result[k] = a[i]
		i++
	}
	for ; j < len(b); k++ {
		result[k] = b[j]
		j++
	}
	return result
}
