package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var max int
	for _, node := range root.Children {
		d := maxDepth(node)
		if max < d {
			max = d
		}
	}
	return max + 1
}
