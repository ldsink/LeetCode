package main

func distributeCandies(candyType []int) int {
	types := make(map[int]bool)
	for _, candy := range candyType {
		types[candy] = true
	}
	max := len(candyType) >> 1
	if max > len(types) {
		max = len(types)
	}
	return max
}
