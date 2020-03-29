package main

func getRow(rowIndex int) []int {
	rowIndex += 1
	if rowIndex < 1 {
		return []int{}
	} else if rowIndex == 1 {
		return []int{1}
	} else if rowIndex == 2 {
		return []int{1, 1}
	}

	var rows [][]int
	for i := 0; i < 2; i++ {
		rows = append(rows, make([]int, rowIndex))
	}

	rows[0][0], rows[0][1] = 1, 1
	for k := 3; k <= rowIndex; k++ {
		curr, prev := k%2, (k+1)%2
		rows[curr][0] = 1
		for i, j := 1, (k-1)/2; i <= j; i++ {
			rows[curr][i] = rows[prev][i] + rows[prev][i-1]
		}
		if k%2 == 1 {
			rows[curr][(k-1)/2] = rows[prev][(k-1)/2-1] * 2
		}
	}

	curr := rowIndex % 2
	for i, j := 0, rowIndex-1; i <= j; i, j = i+1, j-1 {
		rows[curr][j] = rows[curr][i]
	}
	return rows[curr]
}
