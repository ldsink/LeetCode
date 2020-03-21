package main

func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			result = append(result, newInterval)
			result = append(result, intervals[i:]...)
			return result
		} else if intervals[i][1] < newInterval[0] {
			result = append(result, intervals[i])
		} else {
			if newInterval[0] > intervals[i][0] {
				newInterval[0] = intervals[i][0]
			}
			if newInterval[1] < intervals[i][1] {
				newInterval[1] = intervals[i][1]
			}
		}
	}
	return append(result, newInterval)
}
