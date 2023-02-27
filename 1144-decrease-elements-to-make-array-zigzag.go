package main

import "math"

func movesToMakeZigzag(nums []int) int {
	minMoves := math.MaxInt
	for _, less := range []bool{true, false} {
		var moves int
		for i, j, k := 1, less, nums[0]; i < len(nums); i, j = i+1, !j {
			if j {
				if nums[i] >= k {
					moves += (nums[i] - k) + 1
					k -= 1
				} else {
					k = nums[i]
				}
			} else { // bigger
				if nums[i] <= k {
					moves += (k - nums[i]) + 1
				}
				k = nums[i]
			}
		}
		if minMoves > moves {
			minMoves = moves
		}
	}
	return minMoves
}
