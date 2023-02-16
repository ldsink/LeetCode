package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	type Pair struct {
		node  *TreeNode
		level int
	}

	stack := []Pair{{node: root, level: 0}}
	for len(stack) > 0 {
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

	for i := 0; i < len(result); i++ {
		if i%2 == 0 {
			continue
		}
		for j, k := 0, len(result[i])-1; j < k; j, k = j+1, k-1 {
			result[i][j], result[i][k] = result[i][k], result[i][j]
		}
	}
	return result
}
