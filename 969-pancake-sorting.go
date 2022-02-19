package main

func pancakeSort(arr []int) []int {
	var result []int
	for num := len(arr); num > 0; num-- {
		idx := num - 1
		for ; idx >= 0; idx-- {
			if arr[idx] == num {
				break
			}
		}
		if idx == num-1 {
			continue // 这个数字已经排好序了
		}
		// 先把这个转换到到第 1 位，如果已经在第 1 位那就跳过
		if idx != 0 {
			result = append(result, idx+1)
			for i, j := 0, idx; i <= (j >> 1); i++ {
				arr[i], arr[j-i] = arr[j-i], arr[i]
			}
		}
		// 再转换到目标位置
		result = append(result, num)
		for i, j := 0, num-1; i <= (j >> 1); i++ {
			arr[i], arr[j-i] = arr[j-i], arr[i]
		}
	}
	return result
}
