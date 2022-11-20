package main

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestNodes(root *TreeNode, queries []int) [][]int {
	const failed int = -1
	// 把树转换成数组
	var array []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		array = append(array, queue[0].Val)
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
	}
	sort.Ints(array)
	var findMin, findMax func(int) int
	findMin = func(val int) int {
		if val < array[0] {
			return failed
		}
		start, end := 0, len(array)-1
		for start != end {
			middle := (start + end) >> 1
			if array[middle+1] <= val {
				start = middle + 1
			} else {
				end = middle
			}
		}
		return array[start]
	}
	findMax = func(val int) int {
		if val > array[len(array)-1] {
			return failed
		}
		start, end := 0, len(array)-1
		for start != end {
			middle := (start + end) >> 1
			if array[middle] >= val {
				end = middle
			} else {
				start = middle + 1
			}
		}
		return array[start]
	}
	var result [][]int
	for _, query := range queries {
		result = append(result, []int{findMin(query), findMax(query)})
	}
	return result
}
