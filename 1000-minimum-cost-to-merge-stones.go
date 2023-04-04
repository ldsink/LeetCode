package main

import (
	"math"
)

func mergeStones(stones []int, k int) int {
	if k > 2 { // 提前处理始终无法合并情况
		if len(stones)%(k-1) != 1 {
			return -1
		}
	}
	// 缓存下石头总数，方便后续快速求某个范围内的石头
	sum := []int{0}
	for i, stone := range stones {
		sum = append(sum, sum[i]+stone)
	}
	// 每次合并两堆的情况下，简化版动规
	if k == 2 {
		minCost := make(map[int]int)
		var f func(int, int) int
		f = func(start int, end int) int {
			if start == end {
				return 0
			}
			key := start<<10 ^ end
			if c, ok := minCost[key]; ok {
				return c
			}
			if start+1 == end {
				minCost[key] = stones[start] + stones[end]
				return minCost[key]
			}
			cost := math.MaxInt
			for i := start; i < end; i++ {
				if c := f(start, i) + f(i+1, end); cost > c {
					cost = c
				}
			}
			minCost[key] = cost + sum[end+1] - sum[start]
			return minCost[key]
		}
		return f(0, len(stones)-1)
	}
	// 普通的动规
	minCost := make(map[int]int)
	queue := []int{1<<len(stones) - 1}
	// 每个位置上有 1 表示有一堆有效的石头
	for len(queue) > 0 {
		status := queue[0]
		queue = queue[1:]
		// 先还原当前的状态
		var current []int
		for i, j := 0, status; i < len(stones); i, j = i+1, j>>1 {
			if j&1 == 1 {
				current = append(current, i)
			}
		}
		var prefix, suffix int
		for i := k; i < len(current); i++ {
			suffix ^= 1 << current[i]
		}
		current = append(current, len(stones)) // 最后一位放个标记
		for i := 0; i+k < len(current); i++ {
			prefix ^= 1 << current[i]
			cost := minCost[status] + sum[current[i+k]] - sum[current[i]]
			next := prefix ^ suffix // 由当前状态扩展出的下一个状态
			if v, ok := minCost[next]; ok {
				if v > cost {
					minCost[next] = cost
				}
			} else {
				minCost[next] = cost
				if next != 1 {
					queue = append(queue, next)
				}
			}
			suffix ^= 1 << current[i+k] // 尾部移除一个石堆
		}
		delete(minCost, status)
	}
	return minCost[1]
}
