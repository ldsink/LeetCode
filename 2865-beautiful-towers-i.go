package main

func maximumSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	left, right, used := make([]int64, n), make([]int64, n), make([]int, n)
	left[0], used[0] = int64(maxHeights[0]), maxHeights[0]
	for i := 1; i < n; i++ {
		left[i] = left[i-1] + int64(maxHeights[i])
		used[i] = maxHeights[i]
		for j := i - 1; j >= 0 && used[j] > used[i]; j-- {
			left[i] -= int64(used[j] - used[i])
			used[j] = used[i]
		}
	}
	right[n-1], used[n-1] = int64(maxHeights[n-1]), maxHeights[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] + int64(maxHeights[i])
		used[i] = maxHeights[i]
		for j := i + 1; j < n && used[j] > used[i]; j++ {
			right[i] -= int64(used[j] - used[i])
			used[j] = used[i]
		}
	}
	var result int64
	for i := 0; i < n; i++ {
		if j := left[i] + right[i] - int64(maxHeights[i]); result < j {
			result = j
		}
	}
	return result
}
