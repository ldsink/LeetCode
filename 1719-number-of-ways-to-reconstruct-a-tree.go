package main

import (
	"math"
)

func checkWays(pairs [][]int) int {
	linked := make(map[int]map[int]bool)
	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		if linked[a] == nil {
			linked[a] = make(map[int]bool)
		}
		if linked[b] == nil {
			linked[b] = make(map[int]bool)
		}
		linked[a][b], linked[b][a] = true, true
	}

	// 检测是否存在根节点
	var maxEdgeCount int
	for _, edges := range linked {
		if maxEdgeCount < len(edges) {
			maxEdgeCount = len(edges)
		}
	}
	if maxEdgeCount+1 != len(linked) {
		return 0
	}

	ways := 1
	for _, edges := range linked {
		currDegree := len(edges)
		// 根据出度找到父节点。父节点的连接数一定大于当前节点，并且出度最小的。
		parent := -1
		parentDegree := math.MaxInt
		for node := range edges {
			if len(linked[node]) >= currDegree && parentDegree > len(linked[node]) {
				parent = node
				parentDegree = len(linked[node])
			}
		}
		// 找不到父节点就认为失败，并且不是根节点
		if parent == -1 && len(edges) != maxEdgeCount {
			return 0
		}
		// 父节点不是根节点，那需要检查一下包含情况
		if len(edges) != maxEdgeCount {
			// 一个数对 [xi, yi] 出现在 pairs 中 当且仅当 xi 是 yi 的祖先或者 yi 是 xi 的祖先
			for node := range edges {
				if node != parent && !linked[parent][node] {
					return 0
				}
			}
		}
		// 出度相同，这两个节点可以互换
		if parentDegree == currDegree {
			ways = 2
		}
	}
	return ways
}
