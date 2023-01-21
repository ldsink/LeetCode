package main

import "math"

func minSideJumps(obstacles []int) int {
	curr := [3]int{1, 0, 1}
	var prev [3]int
	const unreach = -1
	for i := 1; i < len(obstacles); i++ {
		// 初始化数据
		prev = curr
		curr = [3]int{unreach, unreach, unreach}
		// 可以从前一个点直接跳过来的
		for j := 0; j < 3; j++ {
			if obstacles[i] != j+1 {
				curr[j] = prev[j]
			}
		}
		// 可以在本层级互相跳转
		for j := 0; j < 3; j++ {
			if obstacles[i] == j+1 { // 这个位置上有障碍，跳过
				continue
			}
			for k := 0; k < 3; k++ {
				if curr[k] != unreach && (curr[j] == unreach || curr[j] > curr[k]+1) {
					curr[j] = curr[k] + 1
				}
			}
		}
	}
	result := math.MaxInt
	for i := 0; i < 3; i++ {
		if curr[i] != unreach && result > curr[i] {
			result = curr[i]
		}
	}
	return result
}
