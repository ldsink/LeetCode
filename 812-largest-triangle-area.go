package main

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func largestTriangleArea(points [][]int) float64 {
	var result float64
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				area := float64(abs(points[i][0]*points[j][1]+points[j][0]*points[k][1]+points[k][0]*points[i][1]-points[i][0]*points[k][1]-points[j][0]*points[i][1]-points[k][0]*points[j][1])) / 2
				if result < area {
					result = area
				}
			}
		}
	}
	return result
}
