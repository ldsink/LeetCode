package main

func largestSumOfAverages(nums []int, k int) float64 {
	sums := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	// [i, j) 的平均值
	avg := func(i, j int) float64 {
		return float64(sums[j]-sums[i]) / float64(j-i)
	}
	// 动态规划，只保留当前分组和前一个分组
	var data [2][]float64
	for i := 0; i < 2; i++ {
		data[i] = make([]float64, len(nums))
	}
	for i := 0; i < len(nums); i++ {
		data[0][i] = avg(i, len(nums))
	}
	for i := 1; i < k; i++ {
		curr, prev := i&1, (i^1)&1
		for start := 0; start < len(nums)-i; start++ {
			data[curr][start] = 0
			for end := start + 1; end < len(nums)-i+1; end++ {
				if r := avg(start, end) + data[prev][end]; data[curr][start] < r {
					data[curr][start] = r
				}
			}
		}
	}
	return data[(k-1)&1][0]
}
