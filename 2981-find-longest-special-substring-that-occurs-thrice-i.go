package main

import "sort"

func maximumLength(s string) int {
	data := make(map[rune]map[int]int)
	rs := []rune(s)
	data[rs[len(rs)-1]] = make(map[int]int)
	data[rs[len(rs)-1]][1] = 1
	for i, j := len(s)-2, 1; i >= 0; i-- {
		if rs[i] == rs[i+1] {
			j++
		} else {
			if _, ok := data[rs[i]]; !ok {
				data[rs[i]] = make(map[int]int)
			}
			j = 1
		}
		data[rs[i]][j]++
	}
	result := -1
	type stat struct {
		length, count int
	}
	for _, count := range data { // 遍历每个字符的情况，找出最长的特殊子字符串
		var current []stat
		for length, count := range count {
			current = append(current, stat{length: length, count: count})
		}
		sort.Slice(current, func(i, j int) bool {
			return current[i].length > current[j].length
		})
		var freq int
		for _, s := range current {
			freq += s.count
			if freq >= 3 {
				if result < s.length {
					result = s.length
				}
				break // 满足三次出现，不必再循环
			}
		}
	}
	return result
}
