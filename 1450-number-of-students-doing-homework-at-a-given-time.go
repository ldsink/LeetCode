package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var count int
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			count++
		}
	}
	return count
}
