package main

func candy(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}
	leftLow := make([]int, len(ratings))
	rightLow := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			leftLow[i] = leftLow[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			rightLow[i] = rightLow[i+1] + 1
		}
	}
	var candies int
	for i := 0; i < len(ratings); i++ {
		if leftLow[i] < rightLow[i] {
			candies += 1 + rightLow[i]
		} else {
			candies += 1 + leftLow[i]
		}
	}
	return candies
}
