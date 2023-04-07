package main

import "sort"

func numMovesStonesII(stones []int) []int {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	sort.Ints(stones)
	if stones[len(stones)-1]-stones[0] == len(stones)-1 {
		return []int{0, 0}
	}
	minMoves := len(stones)
	for i, j := 0, 0; i < len(stones); i++ {
		for j+1 < len(stones) && stones[j+1]-stones[i]+1 <= len(stones) {
			j++
		}
		if j-i+1 == len(stones)-1 && stones[j]-stones[i]+1 == len(stones)-1 {
			minMoves = min(minMoves, 2)
		} else {
			minMoves = min(minMoves, len(stones)-(j-i+1))
		}
	}
	maxMoves := max(stones[len(stones)-2]-stones[0]+1, stones[len(stones)-1]-stones[1]+1) - (len(stones) - 1)
	return []int{minMoves, maxMoves}
}
