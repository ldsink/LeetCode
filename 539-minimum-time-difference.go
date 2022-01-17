package main

import (
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	var times []int
	for _, str := range timePoints {
		hour, _ := strconv.Atoi(str[:2])
		minute, _ := strconv.Atoi(str[3:])
		times = append(times, hour*60+minute)
	}
	sort.Ints(times)
	total := 60 * 24
	max := total
	for i := 0; i < len(times); i++ {
		j := (i + 1) % len(times)
		k := (times[j] + total - times[i]) % total
		if max > k {
			max = k
		}
	}
	return max
}
