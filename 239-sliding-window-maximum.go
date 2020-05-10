package main

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSlidingWindow(nums []int, k int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	for i := 0; i < len(nums); i += k {
		l := min(len(nums), i+k) - 1
		left[i] = nums[i]
		for j := i + 1; j <= l; j++ {
			left[j] = max(nums[j], left[j-1])
		}
		right[l] = nums[l]
		for j := l - 1; j >= i; j-- {
			right[j] = max(nums[j], right[j+1])
		}
	}

	result := make([]int, len(nums)-k+1)
	for i := 0; i < len(result); i++ {
		l := min(i+k, len(nums)) - 1
		result[i] = max(right[i], left[l])
	}
	return result
}
