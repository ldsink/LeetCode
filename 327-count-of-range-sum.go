package main

type segNode struct {
	lower, upper, count int
	left, right         *segNode
}

func insert(val int, node *segNode) {
	node.count++
	if node.lower == node.upper {
		return
	}
	middle := (node.lower + node.upper) >> 1
	if val <= middle {
		if node.left == nil {
			node.left = &segNode{lower: node.lower, upper: middle}
		}
		insert(val, node.left)
		return
	}
	if node.right == nil {
		node.right = &segNode{lower: middle + 1, upper: node.upper}
	}
	insert(val, node.right)
}

func getCount(val int, node *segNode) int {
	if val < node.lower {
		return 0
	} else if node.upper <= val {
		return node.count
	}
	middle := (node.lower + node.upper) >> 1
	if val <= middle {
		if node.left == nil {
			return 0
		}
		return getCount(val, node.left)
	}
	var count int
	if node.left != nil {
		count += node.left.count
	}
	if node.right != nil {
		count += getCount(val, node.right)
	}
	return count
}

func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}

	min := nums[0]
	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if min > sum {
			min = sum
		}
		if max < sum {
			max = sum
		}
	}
	if min > 0 {
		min = 0
	}
	if max < 0 {
		max = 0
	}
	root := &segNode{lower: min, upper: max}
	insert(0, root)

	sum = 0
	var count int
	for _, num := range nums {
		sum += num
		count += getCount(sum-lower, root) - getCount(sum-upper-1, root)
		insert(sum, root)
	}
	return count
}
