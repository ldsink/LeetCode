package main

import (
	"fmt"
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	var result []string
	if root.Left != nil {
		r := binaryTreePaths(root.Left)
		for _, p := range r {
			result = append(result, fmt.Sprintf("%d->%s", root.Val, p))
		}
	}
	if root.Right != nil {
		r := binaryTreePaths(root.Right)
		for _, p := range r {
			result = append(result, fmt.Sprintf("%d->%s", root.Val, p))
		}
	}
	return result
}
