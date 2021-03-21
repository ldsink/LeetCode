package main

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	start := 2
	end := num
	for start < end {
		middle := (start + end) >> 1
		if middle*middle < num {
			start = middle + 1
		} else {
			end = middle
		}
	}
	return start*start == num
}
