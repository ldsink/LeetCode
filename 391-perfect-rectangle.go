package main

func isRectangleCover(rectangles [][]int) bool {
	xMin, yMin, xMax, yMax := 100000, 100000, -100000, -100000
	var sumArea int
	for _, rectangle := range rectangles {
		if rectangle[0] < xMin {
			xMin = rectangle[0]
		}
		if rectangle[1] < yMin {
			yMin = rectangle[1]
		}
		if rectangle[2] > xMax {
			xMax = rectangle[2]
		}
		if rectangle[3] > yMax {
			yMax = rectangle[3]
		}
		sumArea += (rectangle[2] - rectangle[0]) * (rectangle[3] - rectangle[1])
	}

	if sumArea != (xMax-xMin)*(yMax-yMin) {
		return false
	}

	visited := make(map[[2]int]bool)
	for i := 0; i < len(rectangles); i++ {
		points := [][2]int{
			{rectangles[i][0], rectangles[i][1]},
			{rectangles[i][0], rectangles[i][3]},
			{rectangles[i][2], rectangles[i][1]},
			{rectangles[i][2], rectangles[i][3]},
		}
		for _, point := range points {
			if visited[point] {
				delete(visited, point)
			} else {
				visited[point] = true
			}
		}
	}
	if len(visited) != 4 {
		return false
	}
	points := [][2]int{
		{xMin, yMin},
		{xMin, yMax},
		{xMax, yMin},
		{xMax, yMax},
	}
	for _, point := range points {
		if !visited[point] {
			return false
		}
	}
	return true
}
