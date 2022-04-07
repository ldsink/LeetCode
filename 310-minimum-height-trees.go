package main

import "sort"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n <= 2 {
		var result []int
		for i := 0; i < n; i++ {
			result = append(result, i)
		}
		return result
	}

	adjacency := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		adjacency[i] = make(map[int]bool)
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adjacency[a][b] = true
		adjacency[b][a] = true
	}

	var leaves []int
	count := n
	for i := 0; i < n; i++ {
		if len(adjacency[i]) == 1 {
			leaves = append(leaves, i)
		}
	}
	for len(leaves) < count {
		var newLeaves []int
		for i := 0; i < len(leaves); i++ {
			n := leaves[i]
			for k := range adjacency[n] {
				delete(adjacency[k], n)
				if len(adjacency[k]) == 1 {
					newLeaves = append(newLeaves, k)
				}
			}
			adjacency[n] = nil
		}
		count -= len(leaves)
		leaves = newLeaves
	}
	sort.Ints(leaves)
	return leaves
}
