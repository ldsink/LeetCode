package main

type segTreeNode struct {
	left, right  int
	count        int
	lNode, rNode *segTreeNode
}

func find(left, right int, root *segTreeNode) int {
	if root == nil {
		return 0
	}
	if left <= root.left && root.right <= right {
		return root.count
	}
	middle := (root.left + root.right) >> 1
	if right < middle {
		return find(left, right, root.lNode)
	} else if left > middle {
		return find(left, right, root.rNode)
	}
	r := find(left, middle, root.lNode)
	if middle+1 <= right {
		r += find(middle+1, right, root.rNode)
	}
	return r
}

func update(num int, root *segTreeNode) {
	root.count++
	if root.left == root.right {
		return
	}
	middle := (root.left + root.right) >> 1
	if num <= middle {
		if root.lNode == nil {
			root.lNode = &segTreeNode{left: root.left, right: middle}
		}
		update(num, root.lNode)
	} else {
		if root.rNode == nil {
			root.rNode = &segTreeNode{left: middle + 1, right: root.right}
		}
		update(num, root.rNode)
	}
}

func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	// initialize
	max, min := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
		if min > nums[i] {
			min = nums[i]
		}
	}
	root := segTreeNode{left: min, right: max}

	// find and update
	result := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == min {
			result[i] = 0
		} else {
			result[i] = find(min, nums[i]-1, &root)
		}
		update(nums[i], &root)
	}
	return result
}
