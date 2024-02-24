package main

func closestNodes(root *TreeNode, queries []int) [][]int {
	// Convert tree to array
	var array []int
	var add func(*TreeNode)
	add = func(node *TreeNode) {
		if node.Left != nil {
			add(node.Left)
		}
		array = append(array, node.Val)
		if node.Right != nil {
			add(node.Right)
		}
	}
	add(root)
	// find min index which array[index] >= value
	findIndex := func(v int) int {
		start, end := 0, len(array)-1
		for start != end {
			middle := (start + end) >> 1
			if array[middle] >= v {
				end = middle
			} else {
				start = middle + 1
			}
		}
		return start
	}
	// Get result
	var result [][]int
	for _, query := range queries {
		if query > array[len(array)-1] {
			result = append(result, []int{array[len(array)-1], -1})
		} else if query < array[0] {
			result = append(result, []int{-1, array[0]})
		} else {
			idx := findIndex(query)
			if array[idx] == query {
				result = append(result, []int{array[idx], array[idx]})
			} else {
				result = append(result, []int{array[idx-1], array[idx]})
			}
		}
	}
	return result
}
