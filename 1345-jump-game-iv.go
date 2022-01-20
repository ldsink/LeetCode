package main

func minJumps(arr []int) int {
	// 每个位置的最小跳动次数。0 代表未处理，1 是起点。jumps[n] - 1 为实际次数。
	jumps := make([]int, len(arr))
	// 每个数值出现的索引位置数
	indexes := make(map[int][]int)
	for idx, num := range arr {
		indexes[num] = append(indexes[num], idx)
	}
	// 广度优先搜索
	jumps[0] = 1
	queue := []int{0}
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		// i + 1 满足：i + 1 < arr.length
		if i+1 < len(arr) && jumps[i+1] == 0 {
			jumps[i+1] = jumps[i] + 1
			queue = append(queue, i+1)
		}
		// i - 1 满足：i - 1 >= 0
		if i-1 >= 0 && jumps[i-1] == 0 {
			jumps[i-1] = jumps[i] + 1
			queue = append(queue, i-1)
		}
		// arr[i] == arr[j] 且 i != j
		if len(indexes[arr[i]]) > 0 {
			for _, idx := range indexes[arr[i]] {
				if jumps[idx] != 0 {
					continue
				}
				jumps[idx] = jumps[i] + 1
				queue = append(queue, idx)
			}
			delete(indexes, arr[i]) // 第一次更新就更新完了，以后肯定没必要再更新
		}
	}
	// 从 1 开始，所以需要减去 1 才是正确结果
	return jumps[len(jumps)-1] - 1
}
