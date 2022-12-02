package main

func minOperations(boxes string) []int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var balls, result []int
	for i, b := range boxes {
		if b == '1' {
			balls = append(balls, i)
		}
	}
	for i := 0; i < len(boxes); i++ {
		var moves int
		for _, j := range balls {
			moves += abs(j - i)
		}
		result = append(result, moves)
	}
	return result
}
