package main

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	var stack []int
	result := make([]int, n)
	start := make([]int, n)
	for _, log := range logs {
		parts := strings.SplitN(log, ":", 3)
		functionId, _ := strconv.Atoi(parts[0])
		timestamp, _ := strconv.Atoi(parts[2])
		if parts[1] == "start" {
			if len(stack) != 0 {
				prev := stack[len(stack)-1]
				if prev != functionId {
					result[stack[len(stack)-1]] += timestamp - start[stack[len(stack)-1]]
					start[functionId] = timestamp
				}
			} else {
				start[functionId] = timestamp
			}
			stack = append(stack, functionId)
		} else {
			result[functionId] += timestamp - start[functionId] + 1
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				start[stack[len(stack)-1]] = timestamp + 1
			}
		}
	}
	return result
}
