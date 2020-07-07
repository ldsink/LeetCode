package main

func lengthOfLIS(nums []int) int {
	var min []int
	for _, n := range nums {
		if len(min) == 0 {
			min = append(min, n)
			continue
		}
		for i := len(min) - 1; i >= 0; i-- {
			if n <= min[i] {
				continue
			}
			if j := i + 1; j == len(min) {
				min = append(min, n)
			} else if min[j] > n {
				min[j] = n
			}
		}
		if n < min[0] {
			min[0] = n
		}
	}
	return len(min)
}
