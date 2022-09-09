package main

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	nums := make([]int, n+1)
	nums[0], nums[1] = 0, 1
	result := 1
	for i := 2; i <= n; i++ {
		if i&1 == 0 {
			nums[i] = nums[i>>1]
		} else {
			nums[i] = nums[i>>1] + nums[i>>1+1]
		}
		if result < nums[i] {
			result = nums[i]
		}
	}
	return result
}
