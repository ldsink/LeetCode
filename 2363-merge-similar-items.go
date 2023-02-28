package main

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	bag := make(map[int]int)
	for _, item := range items1 {
		bag[item[0]] += item[1]
	}
	for _, item := range items2 {
		bag[item[0]] += item[1]
	}
	var result [][]int
	for k, v := range bag {
		result = append(result, []int{k, v})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}
