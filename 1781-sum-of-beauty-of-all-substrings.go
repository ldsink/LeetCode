package main

import "math"

func beautySum(s string) int {
	const L = 26
	const A = 'a'
	getBeauty := func(start int) (sum int) {
		count := make([]int, L)
		var max, min int
		for end := start; end < len(s); end++ {
			idx := s[end] - A
			count[idx]++
			if max < count[idx] {
				max = count[idx]
			}
			if min > count[idx] { // 新加入的就是最小的
				min = count[idx]
			} else if min == count[idx]-1 { // 这次更新可能造成 min 变化，重新计算一下 min
				min = math.MaxInt
				for i := 0; i < L; i++ {
					if count[i] != 0 && min > count[i] {
						min = count[i]
					}
				}
			}
			sum += max - min
		}
		return
	}
	var result int
	for i := 0; i < len(s); i++ {
		result += getBeauty(i)
	}
	return result
}
