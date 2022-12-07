package main

func minOperations(nums1 []int, nums2 []int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var val1, val2 [7]int
	for _, num := range nums1 {
		val1[num]++
	}
	for _, num := range nums2 {
		val2[num]++
	}
	var sum1, sum2 int
	for i := 1; i < 7; i++ {
		sum1 += i * val1[i]
		sum2 += i * val2[i]
	}
	// 始终保持 sum1 <= sum2，逐步进行调整
	if sum1 > sum2 {
		sum1, sum2 = sum2, sum1
		val1, val2 = val2, val1
	}
	// 尝试进行一次调整
	adjust := func() bool {
		for val := min(5, sum2-sum1); val > 0; val-- {
			// 尝试从小的数组增加一次值
			for i := 1; i+val < 7; i++ {
				if val1[i] > 0 {
					val1[i]--
					val1[i+val]++
					sum1 += val
					return true
				}
			}
			// 尝试从大的数字减少一次值
			for i := 6; i-val > 0; i-- {
				if val2[i] > 0 {
					val2[i]--
					val2[i-val]++
					sum2 -= val
					return true
				}
			}
		}
		return false
	}
	var count int
	for ; sum1 < sum2; count++ {
		if !adjust() {
			return -1
		}
	}
	return count
}
