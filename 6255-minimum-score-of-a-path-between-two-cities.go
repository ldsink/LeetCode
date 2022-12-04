package main

import (
	"math"
)

func minScore(n int, roads [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	distance := make(map[[2]int]int)
	saveDistance := func(a, b, d int) {
		if a > b {
			a, b = b, a
		}
		distance[[2]int{a, b}] = d
	}
	getDistance := func(a, b int) int {
		if a > b {
			a, b = b, a
		}
		return distance[[2]int{a, b}]
	}
	linked := make(map[int][]int)
	for _, road := range roads {
		a, b, d := road[0], road[1], road[2]
		saveDistance(a, b, d)
		linked[a] = append(linked[a], b)
		linked[b] = append(linked[b], a)
	}
	scores := make([]int, n+1)
	for i := 1; i <= n; i++ {
		scores[i] = math.MaxInt
	}
	inQueue := make([]bool, n+1)
	for queue := []int{1}; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		inQueue[curr] = false
		for _, next := range linked[curr] {
			if d := min(scores[curr], getDistance(curr, next)); scores[next] > d {
				scores[next] = d
				if !inQueue[next] {
					inQueue[next] = true
					queue = append(queue, next)
				}
			}
		}
	}
	return scores[n]
}
