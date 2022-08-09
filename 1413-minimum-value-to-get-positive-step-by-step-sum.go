package main

func minStartValue(nums []int) int {
	result, sum := 1, 1
	for _, num := range nums {
		sum += num
		if sum < 1 {
			delta := 1 - sum
			result += delta
			sum += delta
		}
	}
	return result
}
