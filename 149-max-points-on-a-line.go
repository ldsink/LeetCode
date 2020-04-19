package main

import (
	"fmt"
	"math"
)

func getLine(p1, p2 [2]int) [2]float64 {
	x1, y1, x2, y2 := p1[0], p1[1], p2[0], p2[1]
	if k := float64(y1-y2) / float64(x1-x2); math.IsInf(k, -1) {
		return [2]float64{math.Inf(-1), float64(x1)}
	} else {
		return [2]float64{k, float64(y1) - k*float64(x1)}
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
	processed := make(map[[2]float64]bool)
	for i := 0; i < len(uniquePoints); i++ {
		count := make(map[[2]float64]int)
		for j := i + 1; j < len(uniquePoints); j++ {
			key := getLine(uniquePoints[i], uniquePoints[j])
			fmt.Println(key)
			if _, ok := processed[key]; ok {
				continue
			}
			count[key] += uniquePointsCount[uniquePoints[j]]
			fmt.Println(count[key], len(count))
		}
		for key, value := range count {
			fmt.Println("current", key, value)
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
	fmt.Println(r)
}
