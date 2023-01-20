package main

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	count := make(map[int]map[int]bool)
	for _, log := range logs {
		user, time := log[0], log[1]
		if count[user] == nil {
			count[user] = make(map[int]bool)
		}
		count[user][time] = true
	}
	result := make([]int, k)
	for _, times := range count {
		result[len(times)-1]++
	}
	return result
}
