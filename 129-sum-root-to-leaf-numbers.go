package main

import "strconv"

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	type weight struct {
		node *TreeNode
		num  string
	}
	stack := []weight{{node: root, num: ""}}
	var result int
	for i := 0; i < len(stack); i++ {
		s := stack[i].num + strconv.Itoa(stack[i].node.Val)
		if stack[i].node.Left == nil && stack[i].node.Right == nil {
			v, _ := strconv.Atoi(s)
			result += v
			continue
		}
		if stack[i].node.Left != nil {
			stack = append(stack, weight{node: stack[i].node.Left, num: s})
		}
		if stack[i].node.Right != nil {
			stack = append(stack, weight{node: stack[i].node.Right, num: s})
		}
	}
	return result
}
