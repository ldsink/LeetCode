package main

func distributeCandies(n int, limit int) int {
	var result int
	for i := 0; i <= limit && i <= n; i++ {
		for j := 0; j <= limit && i+j <= n; j++ {
			if i+j+limit >= n {
				result++
			}
		}
	}
	return result
}
