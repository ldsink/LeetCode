package main

func maximalNetworkRank(n int, roads [][]int) int {
	count := make([]int, n)
	var linked [][]int
	for i := 0; i < n; i++ {
		linked = append(linked, make([]int, n))
	}
	for _, road := range roads {
		count[road[0]]++
		count[road[1]]++
		linked[road[0]][road[1]] = 1
		linked[road[1]][road[0]] = 1
	}
	var result int
	for a := 0; a < n; a++ {
		for b := a + 1; b < n; b++ {
			if c := count[a] + count[b] - linked[a][b]; result < c {
				result = c
			}
		}
	}
	return result
}
