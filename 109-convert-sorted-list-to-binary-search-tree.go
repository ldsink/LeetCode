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

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var nums []int
	for ; head != nil; head = (*head).Next {
		nums = append(nums, (*head).Val)
	}
	return sortedArrayToBST(nums)
}
