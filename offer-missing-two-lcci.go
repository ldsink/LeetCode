package main

func missingTwo(nums []int) []int {
	n := len(nums) + 2
	sum := (n * (n + 1)) >> 1
	for _, num := range nums {
		sum -= num
	}
	middle := sum >> 1
	oneSum := (middle * (middle + 1)) >> 1
	for _, num := range nums {
		if num <= middle {
			oneSum -= num
		}
	}
	return []int{oneSum, sum - oneSum}
}
