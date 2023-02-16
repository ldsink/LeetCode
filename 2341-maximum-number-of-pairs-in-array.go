package main

func numberOfPairs(nums []int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	result := [2]int{}
	for _, num := range count {
		result[0] += num >> 1
		result[1] += num & 1
	}
	return result[:]
}
