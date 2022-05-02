package main

import (
	"regexp"
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	digitLog := regexp.MustCompile("^\\w+(\\s\\d+)+$")
	var digitLogs, letterLogs []string
	for _, l := range logs {
		if digitLog.MatchString(l) {
			digitLogs = append(digitLogs, l)
		} else {
			letterLogs = append(letterLogs, l)
		}
	}
	sort.Slice(letterLogs, func(i, j int) bool {
		// The letter-logs are sorted lexicographically by their contents.
		// If their contents are the same, then sort them lexicographically by their identifiers.
		a := strings.SplitN(letterLogs[i], " ", 2)
		b := strings.SplitN(letterLogs[j], " ", 2)
		order := strings.Compare(a[1], b[1])
		return order < 0 || (order == 0 && strings.Compare(a[0], b[0]) < 0)
	})
	return append(letterLogs, digitLogs...)
}
