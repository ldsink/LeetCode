package main

func generate(numRows int) [][]int {
	var result [][]int
	if numRows >= 1 {
		result = append(result, []int{1})
	}
	if numRows >= 2 {
		result = append(result, []int{1, 1})
	}
	for i := 2; i < numRows; i++ {
		line := []int{1}
		for j := 0; j < len(result[i-1])-1; j++ {
			line = append(line, result[i-1][j]+result[i-1][j+1])
		}
		line = append(line, 1)
		result = append(result, line)
	}
	return result
}
