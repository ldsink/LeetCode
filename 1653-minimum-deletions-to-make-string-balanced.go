package main

func minimumDeletions(s string) int {
	var aTotal, bTotal int
	aCount, bCount := make([]int, len(s)), make([]int, len(s))
	for i, c := range s {
		if c == 'a' {
			aTotal++
		} else {
			bTotal++
		}
		aCount[i], bCount[i] = aTotal, bTotal
	}
	result := aTotal
	if result > bTotal {
		result = bTotal
	}
	for i := 0; i < len(s); i++ {
		if r := bCount[i] + aTotal - aCount[i]; result > r {
			result = r
		}
	}
	return result
}
