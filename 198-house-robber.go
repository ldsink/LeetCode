package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	money := make([]int, len(nums))
	money[0] = nums[0]
	if len(nums) > 1 {
		if nums[0] > nums[1] {
			money[1] = nums[0]
		} else {
			money[1] = nums[1]
		}
	}

	for i := 2; i < len(nums); i++ {
		money[i] = money[i-2] + nums[i]
		if money[i] < money[i-1] {
			money[i] = money[i-1]
		}
	}
	return money[len(nums)-1]
}
