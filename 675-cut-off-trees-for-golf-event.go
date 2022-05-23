package main

import "sort"

func cutOffTree(forest [][]int) int {
	m, n := len(forest), len(forest[0])

	type pair struct{ x, y, v int }
	getKey := func(a, b int) int { return a<<32 | b }
	directions := [4][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	bfs := func(ax, ay, bx, by int) int {
		visited := make(map[int]bool)
		visited[getKey(ax, ay)] = true
		queue := []pair{{ax, ay, 0}}
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			if curr.x == bx && curr.y == by {
				return curr.v
			}
			for _, d := range directions {
				x, y := curr.x+d[0], curr.y+d[1]
				if key := getKey(x, y); 0 <= x && x < m && 0 <= y && y < n && !visited[key] && forest[x][y] > 0 {
					visited[key] = true
					queue = append(queue, pair{x, y, curr.v + 1})
				}
			}
		}
		return -1
	}

	var trees []pair
	for i, row := range forest {
		for j, h := range row {
			if h > 1 {
				trees = append(trees, pair{i, j, h})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool { return trees[i].v < trees[j].v })

	result, preX, preY := 0, 0, 0
	for _, tree := range trees {
		steps := bfs(preX, preY, tree.x, tree.y)
		if steps < 0 {
			return -1
		}
		result += steps
		preX, preY = tree.x, tree.y
	}
	return result
}
