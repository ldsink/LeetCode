package main

import "math/bits"

func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	graph := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	var mask, candidate, maxDistance int
	var dfs func(int, int)
	dfs = func(c, d int) {
		mask ^= 1 << c
		if maxDistance < d {
			maxDistance, candidate = d, c
		}
		for _, v := range graph[c] {
			if mask>>v&1 == 1 {
				dfs(v, d+1)
			}
		}
	}
	result := make([]int, n-1)
	result[0] = len(edges) // 距离为 1 的就是 edges 的数量
	for state := 1<<n - 1; state > 0; state-- {
		if bits.OnesCount(uint(state)) >= 3 { // 至少 3 个城市
			mask, maxDistance = state, 0
			dfs(bits.Len(uint(mask))-1, 0) // 找一个点子树的点，验证子树的有效性
			if mask == 0 {
				mask, maxDistance = state, 0
				dfs(candidate, 0)
				result[maxDistance-1]++
			}
		}
	}
	return result
}
