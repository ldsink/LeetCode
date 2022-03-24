package main

func imageSmoother(img [][]int) [][]int {
	m := len(img)
	n := len(img[0])
	var result [][]int
	directions := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for i := 0; i < m; i++ {
		result = append(result, []int{})
		for j := 0; j < n; j++ {
			sum, total := img[i][j], 1
			for _, d := range directions {
				x, y := i+d[0], j+d[1]
				if 0 <= x && x < m && 0 <= y && y < n {
					sum += img[x][y]
					total += 1
				}
			}
			result[i] = append(result[i], sum/total)
		}
	}
	return result
}
