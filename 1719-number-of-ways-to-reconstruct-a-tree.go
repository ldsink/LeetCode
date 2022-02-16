package main

import "fmt"

func getWays(root int, ancestors map[int]bool, linked map[int][]int, edges map[[2]int]bool) int {
	// 当前节点放到祖先节点里面
	ancestors[root] = true
	// 拆分成 N 个合法的子集。
	var subSets [][]int
	inSet := make(map[int]bool)
	ways := 1 // 假设至少可以凑一种
	for _, node := range linked[root] {
		// 祖先节点或者已经加入集合的节点跳过
		if ancestors[node] || inSet[node] {
			continue
		}
		// 基于当前点，划出一个集合
		currentSet := make(map[int]bool)
		currentSet[node], inSet[node] = true, true
		queue := []int{node}
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			for _, a := range linked[node] {
				if ancestors[a] { // 是祖先节点，跳过
					continue
				} else if currentSet[a] { // 已经在当前节点里面了，跳过
					continue
				}
				if inSet[a] { // 不是祖先节点，也不在当前集合。那说明访问到了其他节点的数据，这样成环了，失败，ways 置 0
					ways = 0
					break
				}
				currentSet[a], inSet[a] = true, true
				queue = append(queue, a)
			}
		}
		if ways == 0 { // 当前集合构造失败，停止检查
			break
		}
		var subSet []int
		for node, _ := range currentSet {
			subSet = append(subSet, node)
		}
		subSets = append(subSets, subSet)
	}
	// 前面分组的地方都没出错，那递归检查，每个子组是不是都合法的。
	if ways != 0 {
		if len(subSets) == 0 { // 当前节点就是叶子节点了，那连接数必须合法
			if len(linked[root]) != len(ancestors)-1 { // ancestors 包含自己，所以需要 -1
				ways = 0
			}
		} else { // 看不同子节点的情况
			for _, subSet := range subSets {
				subWays := 0 // 子组的方案数，每个组都需要检查
				// 找到连接数量最多的节点的连接数
				validNum := len(linked[subSet[0]])
				for i := 1; i < len(subSet); i++ {
					if validNum < len(linked[subSet[i]]) {
						validNum = len(linked[subSet[i]])
					}
				}
				for _, r := range subSet {
					if subWays > 1 { // 这一组已经超过 1 种方案，终止
						break
					}
					if len(linked[r]) != validNum { // 子节点只能是当前组连接最多的点
						continue
					}
					subWays = getWays(r, ancestors, linked, edges) // 以 r 为这个子组的根节点
				}
				if subWays == 0 {
					ways = 0
					break
				} else {
					ways *= subWays
					if ways > 1 {
						ways = 2
					}
				}
			}
		}
	}
	delete(ancestors, root)
	return ways
}

func checkWays(pairs [][]int) int {
	// 统计各个点的连接情况
	linked := make(map[int][]int)
	edges := make(map[[2]int]bool)
	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		linked[a] = append(linked[a], b)
		linked[b] = append(linked[b], a)
		edges[[2]int{a, b}] = true
		edges[[2]int{b, a}] = true
	}
	// 合法的根节点肯定是连接数最多的点
	var max int
	var candidates []int
	for _, v := range linked {
		if max < len(v) {
			max = len(v)
		}
	}
	for k, v := range linked {
		if max == len(v) {
			candidates = append(candidates, k)
		}
	}

	var ways int
	ancestors := make(map[int]bool)
	for _, root := range candidates {
		if ways > 1 { // 如果 ways > 1 ，返回 2
			ways = 2
			break
		}
		ways += getWays(root, ancestors, linked, edges) // 尝试基于该节点作为根节点进行构造，增加方法数量
	}
	return ways
}

func main() {
	fmt.Println(checkWays([][]int{{3, 5}, {4, 5}, {2, 5}, {1, 5}, {1, 4}, {2, 4}})) // 1
	fmt.Println(checkWays([][]int{{1, 2}, {2, 3}, {2, 4}, {1, 5}}))                 // 0
}
