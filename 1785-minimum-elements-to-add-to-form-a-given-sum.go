package main

func minElements(nums []int, limit int, goal int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	diff := goal - sum
	if diff == 0 {
		return 0
	} else if diff < 0 {
		diff = -diff
	}
	r := diff / limit
	if diff%limit != 0 {
		r++
	}
	return r
}
