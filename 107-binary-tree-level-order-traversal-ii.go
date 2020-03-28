package main

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	type NodeLevel struct {
		node  *TreeNode
		level int
	}

	stack := []NodeLevel{NodeLevel{node: root, level: 0}}
	for ; len(stack) > 0; {
		nodeLevel := stack[0]
		stack = stack[1:]
		node, level := *(nodeLevel.node), nodeLevel.level
		if len(result) <= level {
			result = append(result, []int{})
		}
		result[level] = append(result[level], node.Val)
		if node.Left != nil {
			stack = append(stack, NodeLevel{node: node.Left, level: level + 1})
		}
		if node.Right != nil {
			stack = append(stack, NodeLevel{node: node.Right, level: level + 1})
		}
	}
	var reversed [][]int
	for i := len(result) - 1; i >= 0; i-- {
		reversed = append(reversed, result[i])
	}
	return reversed
}
