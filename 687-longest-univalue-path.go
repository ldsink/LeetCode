package main

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	type nodeCount struct {
		left, right int
	}
	var getCount func(*TreeNode) nodeCount
	var result int
	getCount = func(node *TreeNode) nodeCount {
		var count nodeCount
		if node.Left != nil {
			leftCount := getCount(node.Left)
			if node.Val == node.Left.Val {
				count.left = 1 + max(leftCount.left, leftCount.right)
			}
		}
		if node.Right != nil {
			rightCount := getCount(node.Right)
			if node.Val == node.Right.Val {
				count.right = 1 + max(rightCount.left, rightCount.right)
			}
		}
		if result < count.left+count.right {
			result = count.left + count.right
		}
		return count
	}
	getCount(root)
	return result
}
