package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	count := make(map[int]int)
	for _, card := range hand {
		count[card]++
	}
	var cards []int
	for card, _ := range count {
		cards = append(cards, card)
	}
	sort.Ints(cards)
	for _, card := range cards {
		if count[card] == 0 {
			continue
		}
		for i, j := 0, count[card]; i < groupSize; i++ {
			if c, ok := count[card+i]; !ok || c < j {
				return false
			}
			count[card+i] -= j
		}
	}
	return true
}
