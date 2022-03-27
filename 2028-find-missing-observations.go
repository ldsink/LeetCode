package main

func missingRolls(rolls []int, mean int, n int) []int {
	// 计算剩余点数总数
	total := mean * (len(rolls) + n)
	for _, roll := range rolls {
		total -= roll
	}
	// 一定无法构建的情况
	if total < n || total > (6*n) {
		return []int{}
	}
	// 分配点数，构造一个符合要求的答案
	result := make([]int, n)
	initRoll := total / n
	for i := 0; i < n; i++ {
		result[i] = initRoll
	}
	total -= n * initRoll
	for i := 0; total > 0 && i < n; i++ {
		max := 6 - result[i]
		if max > total {
			max = total
		}
		result[i] += max
		total -= max
	}
	return result
}
