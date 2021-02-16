package main

func integerBreak(n int) int {
	result := []int{0, 1, 1}
	if n <= 2 {
		return result[n]
	}
	for i := 3; i <= n; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			k := i - j
			left := j
			if left < result[j] {
				left = result[j]
			}
			right := k
			if right < result[k] {
				right = result[k]
			}
			if max < left*right {
				max = left * right
			}
		}
		result = append(result, max)
	}
	return result[n]
}
