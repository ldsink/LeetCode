package main

func averageValue(nums []int) int {
	var sum, count int
	for _, num := range nums {
		if num&1 == 0 && num%3 == 0 {
			sum += num
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / count
}
