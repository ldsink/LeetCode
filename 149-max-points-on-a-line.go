package main

import (
	"fmt"
	"math"
)

func getLine(p1, p2 [2]int) [3]float64 {
	if deltaY, deltaX := p1[1]-p2[1], p1[0]-p2[0]; deltaX == 0 {
		// x = n
		return [3]float64{0, math.Inf(-1), float64(p1[0])}
	} else if deltaY == 0 {
		// y = n
		return [3]float64{0, 0, float64(p1[1])}
	} else {
		for a, b := deltaX, deltaY; ; a, b = b, a%b {
			if a%b == 0 {
				deltaX, deltaY = deltaX/b, deltaY/b
				break
			}
		}
		if deltaX < 0 {
			deltaX, deltaY = -deltaX, -deltaY
		}
		// y = kx + b
		k := float64(deltaY) / float64(deltaX)
		return [3]float64{float64(deltaX), float64(deltaY), float64(p1[1]) - k*float64(p1[0])}
	}
}

func maxPoints(points [][]int) int {
	var uniquePoints [][2]int
	uniquePointsCount := make(map[[2]int]int)
	for i := 0; i < len(points); i++ {
		uniquePointsCount[[2]int{points[i][0], points[i][1]}] += 1
	}
	for point, _ := range uniquePointsCount {
		uniquePoints = append(uniquePoints, point)
	}
	if len(uniquePoints) <= 2 {
		return len(points)
	}

	var result int
	processed := make(map[[3]float64]bool)
	for i := 0; i < len(uniquePoints); i++ {
		count := make(map[[3]float64]int)
		for j := i + 1; j < len(uniquePoints); j++ {
			key := getLine(uniquePoints[i], uniquePoints[j])
			if _, ok := processed[key]; ok {
				continue
			}
			count[key] += uniquePointsCount[uniquePoints[j]]
		}
		for key, value := range count {
			if result < value+uniquePointsCount[uniquePoints[i]] {
				result = value + uniquePointsCount[uniquePoints[i]]
			}
			processed[key] = true
		}
	}
	return result
}

func main() {
	r := maxPoints([][]int{{0, 0}, {94911152, 94911151}, {94911151, 94911150}})
	//r := maxPoints([][]int{{1, 2}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}})
	fmt.Println(r)
}
