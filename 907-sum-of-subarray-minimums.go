package main

func sumSubarrayMins(arr []int) int {
	const mod = 1000000007
	var left, right, result int
	for i := 0; i < len(arr); i++ {
		for left = i - 1; left >= 0 && arr[left] >= arr[i]; left-- {
		}
		left = i - left - 1
		for right = i + 1; right < len(arr) && arr[right] > arr[i]; right++ {
		}
		right = right - i - 1
		result += arr[i] * (left + 1) * (right + 1)
		result %= mod
	}
	return result
}
