package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	var rank []int
	for _, s := range score {
		rank = append(rank, s)
	}
	sort.Ints(rank)
	for i, j := 0, len(rank)/2; i < j; i++ {
		rank[i], rank[len(rank)-i-1] = rank[len(rank)-i-1], rank[i]
	}
	rankMap := make(map[int]string)
	if len(rank) > 0 {
		rankMap[rank[0]] = "Gold Medal"
	}
	if len(rank) > 1 {
		rankMap[rank[1]] = "Silver Medal"
	}
	if len(rank) > 2 {
		rankMap[rank[2]] = "Bronze Medal"
	}
	for i := 3; i < len(rank); i++ {
		rankMap[rank[i]] = strconv.Itoa(i+1)
	}
	var result []string
	for _, s := range score {
		result = append(result, rankMap[s])
	}
	return result
}
