package main

import (
	"strconv"
)

func calPoints(ops []string) int {
	var scores []int
	for _, op := range ops {
		if op == "+" {
			scores = append(scores, scores[len(scores)-1]+scores[len(scores)-2])
		} else if op == "D" {
			scores = append(scores, scores[len(scores)-1]*2)
		} else if op == "C" {
			scores = scores[:len(scores)-1]
		} else {
			score, _ := strconv.Atoi(op)
			scores = append(scores, score)
		}
	}
	var total int
	for _, score := range scores {
		total += score
	}
	return total
}
