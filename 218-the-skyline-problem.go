package main

import "sort"

func getSkyline(buildings [][]int) [][]int {
	height := make(map[int][]int)
	for _, d := range buildings {
		height[d[0]] = append(height[d[0]], d[2])
		height[d[1]] = append(height[d[1]], -d[2])
	}

	points := make([]int, 0, len(height))
	for k := range height {
		points = append(points, k)
	}
	sort.Ints(points)

	var result [][]int
	current := make(map[int]int)
	var i, j int
	for _, x := range points {
		j = i
		for _, h := range height[x] {
			if h > 0 {
				current[h]++
				if j < h {
					j = h
				}
			} else {
				h = -h
				current[h]--
				if current[h] == 0 {
					delete(current, h)
				}
			}
		}
		if len(current) == 0 {
			result = append(result, []int{x, 0})
			i = 0
		} else if i != j {
			result = append(result, []int{x, j})
			i = j
		} else if _, ok := current[i]; !ok {
			i = 0
			for k := range current {
				if i < k {
					i = k
				}
			}
			result = append(result, []int{x, i})
		}
	}
	return result
}
