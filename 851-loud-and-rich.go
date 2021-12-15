package main

func getQuietest(person int, richer map[int][]int, quiet []int, quietest []int) int {
	if quietest[person] != -1 {
		return quietest[person]
	}
	best := person
	for _, father := range richer[person] {
		if f := getQuietest(father, richer, quiet, quietest); quiet[best] > quiet[f] {
			best = f
		}
	}
	quietest[person] = best
	return best
}

func loudAndRich(richer [][]int, quiet []int) []int {
	richerMap := make(map[int][]int)
	for _, pair := range richer {
		a, b := pair[0], pair[1]
		richerMap[b] = append(richerMap[b], a)
	}
	quietest := make([]int, len(quiet))
	for i := 0; i < len(quiet); i++ {
		quietest[i] = -1
	}
	for i := 0; i < len(quiet); i++ {
		getQuietest(i, richerMap, quiet, quietest)
	}
	return quietest
}
