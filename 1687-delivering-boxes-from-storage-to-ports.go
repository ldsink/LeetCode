package main

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	// same 表示不考虑箱子数量，包含该位置需要的趟数
	same := make([]int, len(boxes)+1)
	weights := make([]int, len(boxes)+1)
	for i := 1; i <= len(boxes); i++ {
		if i > 1 {
			same[i] = same[i-1]
			if boxes[i-2][0] != boxes[i-1][0] {
				same[i]++
			}
		}
		weights[i] = weights[i-1] + boxes[i-1][1]
	}

	minBoxIdx := []int{0}
	minTrips := make([]int, len(boxes)+1)
	dp := make([]int, len(boxes)+1)

	for i := 1; i <= len(boxes); i++ {
		for i-minBoxIdx[0] > maxBoxes || weights[i]-weights[minBoxIdx[0]] > maxWeight {
			minBoxIdx = minBoxIdx[1:]
		}

		minTrips[i] = dp[minBoxIdx[0]] + same[i] + 2

		if i != len(boxes) {
			dp[i] = minTrips[i] - same[i+1]
			for len(minBoxIdx) > 0 && dp[i] <= dp[minBoxIdx[len(minBoxIdx)-1]] {
				minBoxIdx = minBoxIdx[:len(minBoxIdx)-1]
			}
			minBoxIdx = append(minBoxIdx, i)
		}
	}

	return minTrips[len(boxes)]
}
