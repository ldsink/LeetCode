package main

func arraySign(nums []int) int {
	var count int
	for _, num := range nums {
		if num < 0 {
			count++
		} else if num == 0 {
			return 0
		}
	}
	if count%2 == 0 {
		return 1
	}
	return -1
}
