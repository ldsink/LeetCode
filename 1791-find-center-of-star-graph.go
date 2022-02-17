package main

func findCenter(edges [][]int) int {
	count := make(map[int]int)
	for _, edge := range edges {
		count[edge[0]]++
		count[edge[1]]++
	}
	for k, v := range count {
		if v == len(count)-1 {
			return k
		}
	}
	return 0
}
