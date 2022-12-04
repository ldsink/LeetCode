package main

import (
	"math"
)

func magnificentSets(n int, edges [][]int) int {
	// 构造图数据
	graph := make(map[int]map[int]bool)
	for i := 1; i <= n; i++ {
		graph[i] = make(map[int]bool)
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a][b] = true
		graph[b][a] = true
	}
	// 染色进行初步分组
	color := make([]int, n+1)         // 每个节点的颜色
	colorNodes := make(map[int][]int) // 每个颜色包含的节点
	for i := 1; i <= n; i++ {
		if color[i] == 0 {
			color[i] = i
			colorNodes[i] = []int{i}
			for queue := []int{i}; len(queue) > 0; queue = queue[1:] {
				for next, _ := range graph[queue[0]] {
					if color[next] == 0 {
						color[next] = i
						colorNodes[i] = append(colorNodes[i], next)
						queue = append(queue, next)
					}
				}
			}
		}
	}
	// 获得以 start 作为起点的最大分组数
	getSetNum := func(start int) int {
		inSet := make([]bool, n+1)
		inSet[start] = true
		var prev, curr []int
		curr = []int{start}
		var count int
		for len(curr) > 0 {
			count++
			// 先检查本层的合法性
			if len(curr) > 1 {
				for i := 0; i < len(curr); i++ {
					for j := i + 1; j < len(curr); j++ {
						if graph[curr[i]][curr[j]] {
							return -1
						}
					}
				}
			}
			// 扩展下一层
			prev, curr = curr, []int{}
			for _, node := range prev {
				for next, _ := range graph[node] {
					if !inSet[next] {
						inSet[next] = true
						curr = append(curr, next)
					}
				}
			}
		}
		return count
	}
	// 通过入度过滤，获得可以作为起点的节点
	getStartNodes := func(nodes []int) (r []int) {
		degree := math.MaxInt
		for _, node := range nodes {
			if degree > len(graph[node]) {
				degree = len(graph[node])
				r = []int{node}
			} else if degree == len(graph[node]) {
				r = append(r, node)
			}
		}
		return
	}
	var m int
	for _, nodes := range colorNodes {
		if len(nodes) == 1 { // 独立的节点，直接分组
			m++
		} else {
			max := -1
			for _, node := range getStartNodes(nodes) {
				if r := getSetNum(node); r != -1 && (max == -1 || max < r) {
					max = r
				}
			}
			if max == -1 {
				return -1
			}
			m += max
		}
	}
	return m
}
