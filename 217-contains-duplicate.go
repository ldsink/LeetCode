package main

func containsDuplicate(nums []int) bool {
	dict := make(map[int]bool)
	for _, num := range nums {
		if dict[num] {
			return true
		}
		dict[num] = true
	}
	return false
}
