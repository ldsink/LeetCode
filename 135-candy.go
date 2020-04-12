package main

func candy(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}
	leftLow := make([]int, len(ratings))
	righLow := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			leftLow[i] = leftLow[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			righLow[i] = righLow[i+1] + 1
		}
	}
	var candies int
	for i := 0; i < len(ratings); i++ {
		if leftLow[i] < righLow[i] {
			candies += 1 + righLow[i]
		} else {
			candies += 1 + leftLow[i]
		}
	}
	return candies
}
