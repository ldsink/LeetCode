package main

func minMaxGame(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for len(nums) != 1 {
		var newNums []int
		for i, j := 0, len(nums)>>1; i < j; i++ {
			if i&1 == 0 {
				newNums = append(newNums, min(nums[i<<1], nums[i<<1+1]))
			} else {
				newNums = append(newNums, max(nums[i<<1], nums[i<<1+1]))
			}
		}
		nums = newNums
	}
	return nums[0]
}
