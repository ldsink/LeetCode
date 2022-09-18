package main

import "sort"

func rectangleArea(rectangles [][]int) int {
	var xAxis, yAxis []int
	for _, rectangle := range rectangles {
		xAxis = append(xAxis, rectangle[0])
		xAxis = append(xAxis, rectangle[2])
		yAxis = append(yAxis, rectangle[1])
		yAxis = append(yAxis, rectangle[3])
	}
	sort.Ints(xAxis)
	sort.Ints(yAxis)

	var covered [][]bool
	for i := 0; i < len(yAxis); i++ {
		covered = append(covered, make([]bool, len(xAxis)))
	}

	getIndex := func(val int, arr []int) int {
		start, end := 0, len(arr)-1
		for start != end {
			middle := (start + end) >> 1
			if val <= arr[middle] {
				end = middle
			} else {
				start = middle + 1
			}
		}
		return start
	}
	var result int
	for _, rectangle := range rectangles {
		xStart, xEnd := getIndex(rectangle[0], xAxis), getIndex(rectangle[2], xAxis)
		yStart, yEnd := getIndex(rectangle[1], yAxis), getIndex(rectangle[3], yAxis)
		for i := yStart; i < yEnd; i++ {
			for j := xStart; j < xEnd; j++ {
				if !covered[i][j] {
					covered[i][j] = true
					result += (yAxis[i+1] - yAxis[i]) * (xAxis[j+1] - xAxis[j])
				}
			}
		}
	}
	return result % 1000000007
}
