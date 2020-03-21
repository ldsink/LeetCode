package main

func canJump(nums []int) bool {
	canDoIt := make([]bool, len(nums))
	if len(canDoIt) > 0 {
		canDoIt[len(nums)-1] = true
	}
	for i := len(nums) - 2; i >= 0; i-- {
		j := i + nums[i]
		if j >= len(nums)-1 {
			canDoIt[i] = true
			continue
		}
		for k := i + 1; k <= j; k++ {
			if canDoIt[k] {
				canDoIt[i] = true
				break
			}
		}
	}
	return canDoIt[0]
}
