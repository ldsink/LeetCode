package main

import (
	"sort"
	"strconv"
)

func alertNames(keyName []string, keyTime []string) []string {
	type Record struct {
		name string
		time int
	}
	var records []Record
	for i := 0; i < len(keyTime); i++ {
		hour, _ := strconv.Atoi(keyTime[i][:2])
		minute, _ := strconv.Atoi(keyTime[i][3:])
		records = append(records, Record{name: keyName[i], time: hour*60 + minute})
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].time < records[j].time
	})
	nameCount := make(map[string]int)
	nameAlert := make(map[string]bool)
	for start, end := 0, 0; end < len(records); end++ {
		for ; start < end && records[end].time-records[start].time > 60; start++ {
			nameCount[records[start].name]--
		}
		nameCount[records[end].name]++
		if nameCount[records[end].name] >= 3 {
			nameAlert[records[end].name] = true
		}
	}
	var result []string
	for name := range nameAlert {
		result = append(result, name)
	}
	sort.Strings(result)
	return result
}
