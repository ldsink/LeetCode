package main

import (
	"math"
)

func maximumSum(arr []int) int {
	result := -math.MaxInt
	var maxSum [2][2]int // 0 代表可删除情况下以当前结尾的最大值，1 代表不可删除情况下以当前结尾的最大值
	maxSum[1][0] = -math.MaxInt / 2
	maxSum[1][1] = -math.MaxInt / 2
	for i := 0; i < len(arr); i++ {
		idx := i % 2
		prev := idx ^ 1
		// 保留删除机会下的最优解
		if maxSum[prev][0]+arr[i] > arr[i] {
			maxSum[idx][0] = maxSum[prev][0] + arr[i]
		} else {
			maxSum[idx][0] = arr[i]
		}
		// 不保留删除际会下的最优解
		maxSum[idx][1] = maxSum[prev][1] + arr[i] // 直接累计之前的队列
		if maxSum[idx][1] < arr[i] {              // 以当前位置开一个新队列
			maxSum[idx][1] = arr[i]
		}
		if maxSum[idx][1] < maxSum[prev][0] {
			maxSum[idx][1] = maxSum[prev][0] // 删除当前位置构建
		}
		// 更新局部最优解
		if result < maxSum[idx][0] {
			result = maxSum[idx][0]
		}
		if result < maxSum[idx][1] {
			result = maxSum[idx][1]
		}
	}
	return result
}
