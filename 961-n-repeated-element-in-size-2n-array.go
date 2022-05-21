package main

func repeatedNTimes(nums []int) int {
	count := make(map[int]bool)
	for _, num := range nums {
		if count[num] {
			return num
		}
		count[num] = true
	}
	return 0
}
