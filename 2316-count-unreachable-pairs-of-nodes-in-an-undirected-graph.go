package main

func countPairs(n int, edges [][]int) int64 {
	linked := make(map[int][]int)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		linked[a] = append(linked[a], b)
		linked[b] = append(linked[b], a)
	}
	color := make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = i
	}
	count := make(map[int]int)
	for i := 0; i < n; i++ {
		if color[i] != i || len(linked[i]) == 0 {
			continue
		}
		count[i]++
		for queue := []int{i}; len(queue) > 0; queue = queue[1:] {
			for _, next := range linked[queue[0]] {
				if next != i && color[next] == next {
					queue = append(queue, next)
					color[next] = i
					count[i]++
				}
			}
		}
	}
	var colorCount []int
	var colorSum int
	for _, c := range count {
		colorCount = append(colorCount, c)
		colorSum += c
	}
	singleNode := n - colorSum

	var result int64
	// 首先算大群组之间的配对数
	for i := 0; i < len(colorCount); i++ {
		for j := i + 1; j < len(colorCount); j++ {
			result += int64(colorCount[i]) * int64(colorCount[j])
		}
	}
	// 再算大群组和单个点之间的配对数
	result += int64(colorSum) * int64(singleNode)
	// 最后加上单个点之间的配对数
	result += int64(singleNode) * int64(singleNode-1) / 2
	return result
}
