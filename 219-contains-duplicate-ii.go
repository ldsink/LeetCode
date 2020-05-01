package main

func containsNearbyDuplicate(nums []int, k int) bool {
	next := make(map[int]int, len(nums))
	for i, n := range nums {
		if next[n] > 0 && i-next[n] < k {
			return true
		}
		next[n] = i + 1
	}
	return false
}
