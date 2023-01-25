package main

import "sort"

func matrixRankTransform(matrix [][]int) [][]int {
	width := len(matrix[0])
	// 求出每个数值在每个单独的集合中的父节点
	parent := make([]int, len(matrix)*width)
	for i := len(matrix)*width - 1; i >= 0; i-- {
		parent[i] = i
	}
	getRoot := func(x int) int {
		if parent[x] == x {
			return x
		}
		root := x
		for ; root != parent[root]; root = parent[root] {
		}
		parent[x] = root
		for x != root {
			x, parent[x] = parent[x], root
		}
		return root
	}
	unionNum := func(root map[int]int, i, j int) {
		if r, ok := root[matrix[i][j]]; ok {
			currRoot := getRoot(i*width + j)
			parent[currRoot] = r
		} else {
			root[matrix[i][j]] = getRoot(i*width + j)
		}
	}
	// 根据行列，把相同的数值设置成一个集合
	for i := 0; i < len(matrix); i++ {
		root := make(map[int]int)
		for j := 0; j < width; j++ {
			unionNum(root, i, j)
		}
	}
	for j := 0; j < width; j++ {
		root := make(map[int]int)
		for i := 0; i < len(matrix); i++ {
			unionNum(root, i, j)
		}
	}
	// 把矩阵转换成图，然后开始搜索排序
	linked := make(map[int][]int)
	sortNodes := func(nodes []int) {
		sort.Slice(nodes, func(i, j int) bool {
			a, b := nodes[i], nodes[j]
			return matrix[a/width][a%width] < matrix[b/width][b%width]
		})
	}
	addEdges := func(nodesMap map[int]bool) { // 把某行/某列的去重节点，排序之后添加到图
		var nodes []int
		for node := range nodesMap {
			nodes = append(nodes, node)
		}
		sortNodes(nodes)
		for i := 0; i < len(nodes)-1; i++ {
			linked[nodes[i]] = append(linked[nodes[i]], nodes[i+1])
		}
	}
	for i := 0; i < len(matrix); i++ {
		nodes := make(map[int]bool)
		for j := 0; j < width; j++ {
			nodes[getRoot(i*width+j)] = true
		}
		addEdges(nodes)
	}
	for j := 0; j < width; j++ {
		nodes := make(map[int]bool)
		for i := 0; i < len(matrix); i++ {
			nodes[getRoot(i*width+j)] = true
		}
		addEdges(nodes)
	}
	// 把图中的边，按照起点从小到大排序，更新秩的值
	rank := make([]int, len(matrix)*width)   // 每个有效位置上的 rank
	for i := 0; i < len(matrix)*width; i++ { // 初始化
		rank[i] = 1
	}
	var nodes []int
	for node := range linked {
		nodes = append(nodes, node)
	}
	sortNodes(nodes)
	for _, node := range nodes {
		for _, next := range linked[node] {
			if rank[next] < rank[node]+1 {
				rank[next] = rank[node] + 1
			}
		}
	}
	// 转换成结果需要的结构
	var result [][]int
	for i := 0; i < len(matrix); i++ {
		result = append(result, make([]int, width))
		for j := 0; j < width; j++ {
			result[i][j] = rank[getRoot(i*width+j)]
		}
	}
	return result
}
