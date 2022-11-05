package main

import "math"

func bestCoordinate(towers [][]int, radius int) []int {
	result := []int{0, 0}
	r := float64(radius)
	for i, max := 0, 0; i <= 50; i++ {
		for j := 0; j <= 50; j++ {
			quality := 0
			for _, tower := range towers {
				if d := math.Sqrt(float64((i-tower[0])*(i-tower[0]) + (j-tower[1])*(j-tower[1]))); d <= r {
					quality += int(float64(tower[2]) / (1 + d))
				}
			}
			if max < quality {
				max, result = quality, []int{i, j}
			}
		}
	}
	return result
}
