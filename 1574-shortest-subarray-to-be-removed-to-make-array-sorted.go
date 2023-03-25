package main

func findLengthOfShortestSubarray(arr []int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	j := len(arr) - 1
	for j > 0 && arr[j-1] <= arr[j] {
		j--
	}
	if j == 0 {
		return 0
	}
	result := j
	for i := 0; i < len(arr); i++ {
		for j < len(arr) && arr[j] < arr[i] {
			j++
		}
		result = min(result, j-i-1)
		if i+1 < len(arr) && arr[i] > arr[i+1] {
			break
		}
	}
	return result
}
