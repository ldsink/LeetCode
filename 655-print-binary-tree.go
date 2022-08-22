package main

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
	var getDepth func(*TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := getDepth(node.Left), getDepth(node.Right)
		if left < right {
			left = right
		}
		return left + 1
	}

	m := getDepth(root)
	n := 1<<m - 1
	var result [][]string
	for i := 0; i < m; i++ {
		var line []string
		for j := 0; j < n; j++ {
			line = append(line, "")
		}
		result = append(result, line)
	}

	var setValue func(*TreeNode, int, int, int)
	setValue = func(node *TreeNode, x, y, width int) {
		result[y][x] = strconv.Itoa(node.Val)
		nextWidth := width >> 1
		if node.Left != nil {
			setValue(node.Left, x-nextWidth+(nextWidth>>1), y+1, nextWidth)
		}
		if node.Right != nil {
			setValue(node.Right, x+nextWidth-(nextWidth>>1), y+1, nextWidth)
		}
	}
	setValue(root, n>>1, 0, n)

	return result
}
