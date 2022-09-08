package main

func constructArray(n int, k int) []int {
	result := []int{1}
	for i, j := 0, k; j > 0; i, j = i+1, j-1 {
		if i&0x1 == 0 {
			result = append(result, result[i]+j)
		} else {
			result = append(result, result[i]-j)
		}
	}
	for i := k + 2; i <= n; i++ {
		result = append(result, i)
	}
	return result
}
