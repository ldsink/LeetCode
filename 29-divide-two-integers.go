package main

func divide(dividend int, divisor int) int {
	result := int64(dividend) / int64(divisor)
	if result < -(1<<31) || result > (1<<31)-1 {
		result = 1<<31 - 1
	}
	return int(result)
}
