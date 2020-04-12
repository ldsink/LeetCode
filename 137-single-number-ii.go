package main

func singleNumber(nums []int) int {
	var one, two int
	for _, num := range nums {
		one = (one ^ num) & (^two)
		two = (^one) & (two ^ num)
	}
	return one ^ two
}
