package main

func gardenNoAdj(n int, paths [][]int) []int {
	linked := make(map[int][]int)
	for _, path := range paths {
		a, b := path[0]-1, path[1]-1
		linked[a] = append(linked[a], b)
		linked[b] = append(linked[b], a)
	}
	result := make([]int, n)
	getColor := func(node int) int {
		for color := 1; color <= 4; color++ {
			success := true
			for _, c := range linked[node] {
				if result[c] == color {
					success = false
					break
				}
			}
			if success {
				return color
			}
		}
		return 0
	}
	for i := 0; i < n; i++ {
		if result[i] != 0 {
			continue
		}
		for queue := []int{i}; len(queue) > 0; queue = queue[1:] {
			node := queue[0]
			result[node] = getColor(node)
			for _, n := range linked[node] {
				if result[n] == 0 {
					queue = append(queue, n)
					result[n] = -1
				}
			}
		}
	}
	return result
}
