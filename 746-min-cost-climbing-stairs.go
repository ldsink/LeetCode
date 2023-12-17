package main

func minCostClimbingStairs(cost []int) int {
	minCost := make([]int, len(cost))
	for i := 2; i < len(cost); i++ {
		minCost[i] = minCost[i-1] + cost[i-1]
		if other := minCost[i-2] + cost[i-2]; minCost[i] > other {
			minCost[i] = other
		}
	}
	idx := len(minCost) - 1
	result := minCost[idx] + cost[idx]
	if other := minCost[idx-1] + cost[idx-1]; result > other {
		result = other
	}
	return result
}
