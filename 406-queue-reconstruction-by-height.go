package main

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0] || (people[i][0] == people[j][0] && people[i][1] > people[j][1])
	})
	indexes := make([]int, len(people))
	result := make([][]int, len(people))
	for i := 0; i < len(indexes); i++ {
		indexes[i] = i
	}
	for _, p := range people {
		result[indexes[p[1]]] = p
		indexes = append(indexes[:p[1]], indexes[p[1]+1:]...)
	}
	return result
}
