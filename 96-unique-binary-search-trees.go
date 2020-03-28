package main

func countByNums(n int, count []int) int {
	if n < 2 {
		return n
	}
	if count[n] > 0 {
		return count[n]
	}

	var result int
	for i := 0; i < n; i++ {
		var left, right int
		if left = countByNums(i, count); left == 0 {
			left = 1
		}
		if right = countByNums(n-i-1, count); right == 0 {
			right = 1
		}
		result += left * right
	}
	count[n] = result
	return result
}

func numTrees(n int) int {
	count := make([]int, n+1)
	return countByNums(n, count)
}
