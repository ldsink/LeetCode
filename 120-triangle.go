package main

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	} else if len(triangle) == 1 {
		return triangle[0][0]
	}

	var minTotal [][]int
	for i := 0; i < 2; i++ {
		minTotal = append(minTotal, make([]int, len(triangle)))
	}
	for i := 0; i < len(triangle); i++ {
		curr, prev := i%2, (i+1)%2
		for j := 0; j < len(triangle[i]); j++ {
			minTotal[curr][j] = triangle[i][j]
			if j-1 >= 0 && j < i {
				if minTotal[prev][j] > minTotal[prev][j-1] {
					minTotal[curr][j] += minTotal[prev][j-1]
				} else {
					minTotal[curr][j] += minTotal[prev][j]
				}
			} else if j-1 >= 0 {
				minTotal[curr][j] += minTotal[prev][j-1]
			} else if j < i {
				minTotal[curr][j] += minTotal[prev][j]
			}
		}
	}
	curr := (len(triangle) - 1) % 2
	result := minTotal[curr][0]
	for i := 1; i < len(triangle); i++ {
		if result > minTotal[curr][i] {
			result = minTotal[curr][i]
		}
	}
	return result
}
