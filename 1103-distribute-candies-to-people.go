package main

func distributeCandies(candies int, num_people int) []int {
	result := make([]int, num_people)
	circle := (num_people * (num_people + 1)) >> 1
	var cost, n int
	for n = 0; cost+n*num_people*num_people+circle <= candies; n++ {
		cost += n*num_people*num_people + circle
	}
	base := (n * (n - 1)) >> 1
	candies -= base*num_people*num_people + circle*n
	for i := 0; i < num_people; i++ {
		result[i] += base*num_people + (i+1)*n
	}
	for i := 0; candies > 0; i++ {
		if candies >= n*num_people+i+1 {
			result[i] += n*num_people + i + 1
			candies -= n*num_people + i + 1
		} else {
			result[i] += candies
			candies = 0
			break
		}
	}
	return result
}
