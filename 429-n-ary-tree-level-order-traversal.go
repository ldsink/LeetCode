package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	type NodeDepth struct {
		node  *Node
		depth int
	}
	queue := []NodeDepth{{node: root, depth: 1}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if len(result) < current.depth {
			result = append(result, []int{})
		}
		result[len(result)-1] = append(result[len(result)-1], current.node.Val)
		for _, node := range current.node.Children {
			queue = append(queue, NodeDepth{node: node, depth: current.depth + 1})
		}
	}
	return result
}
