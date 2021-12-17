package main

func numWaterBottles(numBottles int, numExchange int) int {
	var result, emptyBottles int
	for numBottles > 0 {
		result += numBottles
		emptyBottles += numBottles
		numBottles = emptyBottles / numExchange
		emptyBottles %= numExchange
	}
	return result
}
