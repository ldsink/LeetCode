package main

import "math"

func maximumDetonation(bombs [][]int) int {
	linked := make(map[int][]int)
	for i := 0; i < len(bombs); i++ {
		for j := i + 1; j < len(bombs); j++ {
			x, y := bombs[i][0]-bombs[j][0], bombs[i][1]-bombs[j][1]
			distance := math.Sqrt(float64(x*x + y*y))
			if float64(bombs[i][2]) >= distance {
				linked[i] = append(linked[i], j)
			}
			if float64(bombs[j][2]) >= distance {
				linked[j] = append(linked[j], i)
			}
		}
	}
	var result int
	for i := 0; i < len(bombs); i++ {
		queue, inQueue := []int{i}, make([]bool, len(bombs))
		inQueue[i] = true
		for j := 0; j < len(queue); j++ {
			for _, k := range linked[queue[j]] {
				if inQueue[k] {
					continue
				}
				inQueue[k] = true
				queue = append(queue, k)
			}
		}
		if result < len(queue) {
			result = len(queue)
		}
	}
	return result
}
