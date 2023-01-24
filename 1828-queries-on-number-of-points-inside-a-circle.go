package main

import "math"

func countPoints(points [][]int, queries [][]int) []int {
	var result []int
	for _, query := range queries {
		var count int
		for _, point := range points {
			x := point[0] - query[0]
			y := point[1] - query[1]
			if float64(query[2]) >= math.Sqrt(float64(x*x+y*y)) {
				count++
			}
		}
		result = append(result, count)
	}
	return result
}
