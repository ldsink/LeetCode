package main

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var first, second, firstIndex, secondIndex int
	firstIndex, secondIndex = -1, -1
	for idx, num := range nums {
		if firstIndex == -1 || first < num {
			second, secondIndex = first, firstIndex
			first, firstIndex = num, idx
		} else if secondIndex == -1 || second < num {
			second, secondIndex = num, idx
		}
	}
	if firstIndex != -1 && secondIndex != -1 && first >= 2*second {
		return firstIndex
	}
	return -1
}
