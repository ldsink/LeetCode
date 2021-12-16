package main

func search(nums []int, index, target int, cached map[[2]int]int) int {
	if value, ok := cached[[2]int{index, target}]; ok {
		return value
	}
	cached[[2]int{index, target}] = 0
	if index == len(nums)-1 {
		// 最后一位了，尝试 + 和 - 能不能到 target。+0 和 -0 应该会计算两种
		if nums[index] == target {
			cached[[2]int{index, target}]++
		}
		if nums[index] == -target {
			cached[[2]int{index, target}]++
		}
	} else {
		cached[[2]int{index, target}] = search(nums, index+1, target+nums[index], cached) + search(nums, index+1, target-nums[index], cached)
	}
	return cached[[2]int{index, target}]
}

func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	cached := make(map[[2]int]int)
	return search(nums, 0, target, cached)
}
