package main

import (
	"sort"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	count := make([]int, amount+1)
	for i := 1; i < len(count); i++ {
		count[i] = -1
	}
	sort.Ints(coins)
	for i := len(coins) - 1; i >= 0; i-- {
		for j := coins[i]; j <= amount; j++ {
			if count[j-coins[i]] == -1 {
				continue
			}
			if count[j] == -1 || count[j] > count[j-coins[i]]+1 {
				count[j] = count[j-coins[i]] + 1
			}
		}
	}
	return count[amount]
}
