package main

func minAddToMakeValid(s string) int {
	var result, count int
	for _, r := range s {
		if r == '(' {
			count++
		} else if count > 0 {
			count--
		} else {
			result++
		}
	}
	return result + count
}
