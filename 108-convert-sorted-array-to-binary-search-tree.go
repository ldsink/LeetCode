package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	middle := len(nums) >> 1
	return &TreeNode{Val: nums[middle], Left: sortedArrayToBST(nums[:middle]), Right: sortedArrayToBST(nums[middle+1:])}
}
