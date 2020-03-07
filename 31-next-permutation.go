package main

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		// keep lowest possible order
		for j := i; j < len(nums)-1; j++ {
			if nums[j] <= nums[j+1] {
				break
			}
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
		// try find next greater permutation of numbers
		if i-1 >= 0 && num > nums[i-1] {
			for j := i; j < len(nums); j++ {
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					break
				}
			}
			break
		}
	}
}
