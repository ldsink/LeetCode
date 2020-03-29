package main

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	type Pair struct {
		node  *TreeNode
		level int
	}
	stack := []Pair{{node: root, level: 0}}
	for ; len(stack) > 0; {
		nodeWithLevel := stack[0]
		stack = stack[1:]
		node, level := *(nodeWithLevel.node), nodeWithLevel.level
		if len(result) <= level {
			result = append(result, []int{})
		}
		result[level] = append(result[level], node.Val)
		if node.Left != nil {
			stack = append(stack, Pair{node: node.Left, level: level + 1})
		}
		if node.Right != nil {
			stack = append(stack, Pair{node: node.Right, level: level + 1})
		}
	}
	return result
}
