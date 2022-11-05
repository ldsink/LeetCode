package main

func mergeTriplets(triplets [][]int, target []int) bool {
	var one, two, three bool
	for _, t := range triplets {
		if t[0] <= target[0] && t[1] <= target[1] && t[2] <= target[2] {
			if !one && t[0] == target[0] {
				one = true
			}
			if !two && t[1] == target[1] {
				two = true
			}
			if !three && t[2] == target[2] {
				three = true
			}
			if one && two && three {
				return true
			}
		}
	}
	return false
}
