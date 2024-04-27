package main

import "strconv"

func findColumnWidth(grid [][]int) []int {
	var result []int
	for i := 0; i < len(grid[0]); i++ {
		var l int
		for j := 0; j < len(grid); j++ {
			num := strconv.Itoa(grid[j][i])
			if l < len(num) {
				l = len(num)
			}
		}
		result = append(result, l)
	}
	return result
}
