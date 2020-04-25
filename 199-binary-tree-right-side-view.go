package main

func rightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	type Level struct {
		node  *TreeNode
		level int
	}
	stack := []Level{{node: root, level: 1}}
	for ; len(stack) > 0; stack = stack[1:] {
		if len(result) < stack[0].level {
			result = append(result, stack[0].node.Val)
		} else {
			result[stack[0].level-1] = stack[0].node.Val
		}
		if stack[0].node.Left != nil {
			stack = append(stack, Level{node: stack[0].node.Left, level: stack[0].level + 1})
		}
		if stack[0].node.Right != nil {
			stack = append(stack, Level{node: stack[0].node.Right, level: stack[0].level + 1})
		}
	}
	return result
}
