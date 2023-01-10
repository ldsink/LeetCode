package main

import (
	"math"
)

func crackSafe(n int, k int) string {
	visited := make(map[int]bool)
	highest := int(math.Pow(10, float64(n-1)))
	var result []int
	var dfs func(int)
	dfs = func(node int) {
		for i := 0; i < k; i++ {
			val := node*10 + i
			if !visited[val] {
				visited[val] = true
				dfs(val % highest)
				result = append(result, i)
			}
		}
	}
	dfs(0)
	for i := 1; i < n; i++ {
		result = append(result, 0)
	}
	var rs []rune
	for _, num := range result {
		rs = append(rs, rune('0'+num))
	}
	return string(rs)
}
