package main

import (
	"math"
)

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	linked := make([][]int, n)
	length := make(map[[2]int]int) // 两点之间的距离
	extend := make(map[[2]int]int) // 两点之间可以达到的点
	for _, edge := range edges {
		a, b, d := edge[0], edge[1], edge[2]
		linked[edge[0]] = append(linked[edge[0]], b)
		linked[edge[1]] = append(linked[edge[1]], a)
		length[[2]int{a, b}] = d + 1
		length[[2]int{b, a}] = d + 1
	}
	distance := make([]int, n)
	for i := 1; i < n; i++ {
		distance[i] = math.MaxInt
	}
	inQueue := make([]bool, n)
	for queue := []int{0}; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		inQueue[curr] = false
		for _, next := range linked[curr] {
			if d := distance[curr] + length[[2]int{curr, next}]; d <= maxMoves { // 可以直接走到 next
				extend[[2]int{curr, next}] = length[[2]int{curr, next}] - 1 // 中间的所有点都能走到
				if distance[next] > d {
					distance[next] = d
					if d < maxMoves && !inQueue[next] {
						inQueue[next] = true
						queue = append(queue, next)
					}
				}
			} else if e := maxMoves - distance[curr]; extend[[2]int{curr, next}] < e { // 不能走到，那就尝试更新边的最大长度
				extend[[2]int{curr, next}] = e
			}
		}
	}
	var result int // 最终结果
	// 可以到达的原始节点
	for i := 0; i < n; i++ {
		if distance[i] != math.MaxInt {
			result++
		}
	}
	// 可以到达的新增节点
	for edge := range length {
		a, b := edge[0], edge[1]
		if a > b {
			continue
		}
		result += min(extend[[2]int{a, b}]+extend[[2]int{b, a}], length[[2]int{a, b}]-1)
	}
	return result
}
