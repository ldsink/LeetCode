package main

import "sort"

func fillCups(amount []int) int {
	sort.Ints(amount)
	if amount[0]+amount[1] <= amount[2] {
		return amount[2]
	}
	sum := amount[0] + amount[1] + amount[2]
	return (sum >> 1) + (sum & 1)
}
