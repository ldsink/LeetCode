package main

func nearestValidPoint(x int, y int, points [][]int) int {
	abs := func(x int) int {
		if x >= 0 {
			return x
		}
		return -x
	}
	idx := -1
	var smallest int
	for i, point := range points {
		if point[0] == x || point[1] == y {
			if d := abs(point[0]-x) + abs(point[1]-y); idx == -1 || smallest > d {
				idx, smallest = i, d
			}
		}
	}
	return idx
}
