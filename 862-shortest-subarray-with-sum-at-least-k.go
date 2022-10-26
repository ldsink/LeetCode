package main

func shortestSubarray(nums []int, k int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var start, sum int
	result := -1
	for i := 0; i < len(nums); i++ {
		if sum+nums[i] < 0 {
			start = i + 1
			sum = 0
			continue
		}
		sum += nums[i]
		for j := i - 1; nums[i] < 0 && j >= start; j-- {
			if nums[j] > 0 {
				val := min(nums[j], -nums[i])
				nums[j] -= val
				nums[i] += val
			}
		}
		if sum >= k {
			for ; sum-nums[start] >= k; start++ {
				sum -= nums[start]
			}
			if result == -1 || result > i-start+1 {
				result = i - start + 1
			}
		}
	}
	return result
}
