package main

func maxScore(s string) int {
	var countZero, countOne, leftZero, leftOne, result int
	for _, r := range s {
		if r == '0' {
			countZero++
		} else {
			countOne++
		}
	}
	if s[0] == '0' {
		leftZero++
	} else {
		leftOne++
	}
	for i := 1; i < len(s); i++ {
		if r := leftZero + countOne - leftOne; result < r {
			result = r
		}
		if s[i] == '0' {
			leftZero++
		} else {
			leftOne++
		}
	}
	return result
}
