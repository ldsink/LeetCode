package main

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	var pending, maxProfit, minRotate, revenue, cost int
	minRotate = -1
	for round, customer := range customers {
		pending += customer
		if pending > 4 {
			revenue += 4 * boardingCost
			pending -= 4
		} else {
			revenue += pending * boardingCost
			pending = 0
		}
		cost += runningCost
		if maxProfit < revenue-cost {
			maxProfit = revenue - cost
			minRotate = round + 1
		}
	}
	for round := len(customers); pending > 0; round = round + 1 {
		if pending > 4 {
			revenue += 4 * boardingCost
			pending -= 4
		} else {
			revenue += pending * boardingCost
			pending = 0
		}
		cost += runningCost
		if maxProfit < revenue-cost {
			maxProfit = revenue - cost
			minRotate = round + 1
		}
	}
	return minRotate
}
