package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	var v1, v2 []int
	for _, s := range strings.Split(version1, ".") {
		n, _ := strconv.Atoi(s)
		v1 = append(v1, n)
	}
	for _, s := range strings.Split(version2, ".") {
		n, _ := strconv.Atoi(s)
		v2 = append(v2, n)
	}
	for ; len(v1) < len(v2); {
		v1 = append(v1, 0)
	}
	for ; len(v2) < len(v1); {
		v2 = append(v2, 0)
	}
	for i := 0; i < len(v1); i++ {
		if v1[i] > v2[i] {
			return 1
		} else if v1[i] < v2[i] {
			return -1
		}
	}
	return 0
}
