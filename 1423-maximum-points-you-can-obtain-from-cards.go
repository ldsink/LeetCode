package main

func maxScore(cardPoints []int, k int) int {
	var sum, result, temp int
	for _, p := range cardPoints {
		sum += p
	}
	l := len(cardPoints) - k
	if l == 0 {
		return sum
	}
	for i := 0; i < l; i++ {
		temp += cardPoints[i]
	}
	result = temp
	for s, e := 0, l; e < len(cardPoints); s, e = s+1, e+1 {
		temp += cardPoints[e] - cardPoints[s]
		if result > temp {
			result = temp
		}
	}
	return sum - result
}
