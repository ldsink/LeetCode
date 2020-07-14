package main

import (
	"sort"
)

type ByHeight [][]int

func (a ByHeight) Len() int {
	return len(a)
}

func (a ByHeight) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}

func (a ByHeight) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func reconstructQueue(people [][]int) [][]int {
	sort.Sort(ByHeight(people))
	var result [][]int
	for ; len(people) > 0; {
		for i := 0; i < len(people); i++ {
			k := 0
			for j := 0; j < len(result); j++ {
				if result[j][0] >= people[i][0] {
					k++
				}
			}
			if k == people[i][1] {
				result = append(result, people[i])
				people = append(people[:i], people[i+1:]...)
				break
			}
		}
	}
	return result
}
