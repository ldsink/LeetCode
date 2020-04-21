package main

import (
	"sort"
	"strconv"
	"strings"
)

type Number []string

func (a Number) Len() int {
	return len(a)
}

func (a Number) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Number) Less(i, j int) bool {
	k, l := a[i]+a[j], a[j]+a[i]
	for m := 0; m < len(k); m++ {
		if k[m] > l[m] {
			return true
		} else if k[m] < l[m] {
			return false
		}
	}
	return false
}

func largestNumber(nums []int) string {
	var list []string
	for _, num := range nums {
		list = append(list, strconv.Itoa(num))
	}
	sort.Sort(Number(list))
	result := strings.Join(list, "")
	if strings.HasPrefix(result, "00") {
		return "0"
	}
	return result
}
