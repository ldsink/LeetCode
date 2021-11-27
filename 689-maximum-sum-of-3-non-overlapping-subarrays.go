package main

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	// 预计算连续 k 个项的值
	sum := make([]int, len(nums)-k+1)
	for i := 0; i < k; i++ {
		sum[0] += nums[i]
	}
	for i := 1; i <= len(nums)-k; i++ {
		sum[i] = sum[i-1] - nums[i-1] + nums[i-1+k]
	}

	thirdIdx := make([]int, len(nums)-k+1)
	thirdMax := make([]int, len(nums)-k+1)
	// 计算第三个位置上的最优结果
	for i, j := len(nums)-k, len(nums)-k; i >= 0; i-- {
		thirdIdx[i] = i
		thirdMax[i] = sum[i]
		if i+1 <= j && thirdMax[i] < thirdMax[i+1] {
			thirdMax[i] = thirdMax[i+1]
			thirdIdx[i] = thirdIdx[i+1]
		}
	}
	// 计算中间位置上的最优结果
	secondIdx := make([][2]int, len(nums)-2*k+1)
	secondMax := make([]int, len(nums)-2*k+1)
	for i, j := len(nums)-2*k, len(nums)-2*k; i >= 0; i-- {
		secondIdx[i] = [2]int{i, thirdIdx[i+k]}
		secondMax[i] = sum[i] + thirdMax[i+k]
		if i+1 <= j && secondMax[i] < secondMax[i+1] {
			secondMax[i] = secondMax[i+1]
			secondIdx[i] = secondIdx[i+1]
		}
	}
	// 第一层的结果
	firstIdx := make([][3]int, len(nums)-3*k+1)
	firstMax := make([]int, len(nums)-3*k+1)
	for i, j := len(nums)-3*k, len(nums)-3*k; i >= 0; i-- {
		firstIdx[i] = [3]int{i, secondIdx[i+k][0], secondIdx[i+k][1]}
		firstMax[i] = sum[i] + secondMax[i+k]
		if i+1 <= j && firstMax[i] < firstMax[i+1] {
			firstMax[i] = firstMax[i+1]
			firstIdx[i] = firstIdx[i+1]
		}
	}
	return firstIdx[0][:]
}
