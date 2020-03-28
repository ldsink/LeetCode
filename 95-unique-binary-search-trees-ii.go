package main

func generateByNums(nums []int) []*TreeNode {
	var result []*TreeNode
	if len(nums) == 0 {
		return result
	} else if len(nums) == 1 {
		node := TreeNode{Val: nums[0]}
		result = append(result, &node)
		return result
	}
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		leftNodes := generateByNums(nums[:i])
		if len(leftNodes) == 0 {
			leftNodes = append(leftNodes, nil)
		}
		rightNodes := generateByNums(nums[i+1:])
		if len(rightNodes) == 0 {
			rightNodes = append(rightNodes, nil)
		}
		for _, left := range leftNodes {
			for _, right := range rightNodes {
				r := TreeNode{Val: val, Left: left, Right: right}
				result = append(result, &r)
			}
		}
	}
	return result
}

func generateTrees(n int) []*TreeNode {
	var result []*TreeNode
	if n == 0 {
		return result
	}
	var nums []int
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	return generateByNums(nums)
}
