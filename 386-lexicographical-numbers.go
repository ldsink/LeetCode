package main

import (
	"sort"
	"strconv"
)

func lexicalOrderCheck(n int) []int {
	var s []string
	for i := 1; i <= n; i++ {
		s = append(s, strconv.Itoa(i))
	}
	sort.Strings(s)
	var result []int
	for _, a := range s {
		i, _ := strconv.Atoi(a)
		result = append(result, i)
	}
	return result
}

func extendNums(num, idx int, plus bool, result []int) int {
	if num > len(result) {
		return idx
	}
	result[idx] = num
	idx++

	idx = extendNums(num*10, idx, true, result)

	if plus {
		for i := 1; (num+i)%10 != 0; i++ {
			idx = extendNums(num+i, idx, false, result)
		}
	}
	return idx
}

func lexicalOrder(n int) []int {
	result := make([]int, n)
	extendNums(1, 0, true, result)
	return result
}
