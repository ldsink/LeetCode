package main

func plusOne(digits []int) []int {
	i := len(digits) - 1
	for ; i >= 0; i-- {
		digits[i] += 1
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
	}
	if i < 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
