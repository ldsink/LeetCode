package main

import "sort"

func bts(root *TreeNode) ([]int, []*TreeNode) {
	var nums []int
	var nodes []*TreeNode
	if (*root).Left != nil {
		n, r := bts((*root).Left)
		nums = append(nums, n...)
		nodes = append(nodes, r...)
	}
	nums = append(nums, (*root).Val)
	nodes = append(nodes, root)
	if (*root).Right != nil {
		n, r := bts((*root).Right)
		nums = append(nums, n...)
		nodes = append(nodes, r...)
	}
	return nums, nodes
}

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	nums, nodes := bts(root)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		(*nodes[i]).Val = nums[i]
	}
}
