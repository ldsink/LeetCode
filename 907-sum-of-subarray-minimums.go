package main

func sumSubarrayMins(arr []int) int {
	const mod = 1000000007
	var left, right, result int
	for i := 0; i < len(arr); i++ {
		if i == 0 || arr[i] > arr[i-1] {
			left = i - 1
		}
		for ; left >= 0 && arr[left] >= arr[i]; left-- {
		}
		for right = i + 1; right < len(arr) && arr[right] > arr[i]; right++ {
		}
		result += arr[i] * (i - left) * (right - i)
		result %= mod
	}
	return result
}
