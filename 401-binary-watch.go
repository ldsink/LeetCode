package main

import (
	"math/bits"
	"strconv"
)

func getMap(end int) map[int][]int {
	result := make(map[int][]int)
	for i := 0; i < end; i++ {
		n := bits.OnesCount(uint(i))
		result[n] = append(result[n], i)
	}
	return result
}

func readBinaryWatch(turnedOn int) []string {
	var result []string
	if turnedOn >= 9 {
		return result
	}
	hours := getMap(12)
	minutes := getMap(60)
	for i := 0; i < 4 && i <= turnedOn; i++ {
		j := turnedOn - i
		if j < 0 || j > 6 {
			continue
		}
		if len(hours[i]) == 0 {
			continue
		}
		if len(minutes[j]) == 0 {
			continue
		}
		// 拼装成字符串，加到结果数组
		for _, hour := range hours[i] {
			for _, minute := range minutes[j] {
				m := strconv.Itoa(minute)
				if minute < 10 {
					m = "0" + m
				}
				result = append(result, strconv.Itoa(hour)+":"+m)
			}
		}
	}
	return result
}
