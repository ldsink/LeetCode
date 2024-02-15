package main

import (
	"sort"
)

func verticalTraversal(root *TreeNode) [][]int {
	data := make(map[int][][]int)
	var search func(*TreeNode, int, int)
	search = func(node *TreeNode, col, row int) {
		if _, ok := data[col]; !ok {
			data[col] = [][]int{}
		}
		for len(data[col]) <= row {
			data[col] = append(data[col], []int{})
		}
		data[col][row] = append(data[col][row], node.Val)
		if node.Left != nil {
			search(node.Left, col-1, row+1)
		}
		if node.Right != nil {
			search(node.Right, col+1, row+1)
		}
	}
	search(root, 0, 0)
	var cols []int
	for col := range data {
		cols = append(cols, col)
	}
	sort.Ints(cols)
	var result [][]int
	for idx, col := range cols {
		result = append(result, []int{})
		for _, d := range data[col] {
			sort.Ints(d)
			result[idx] = append(result[idx], d...)
		}
	}
	return result
}
