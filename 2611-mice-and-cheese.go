package main

import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	var array []int
	for i := 0; i < len(reward1); i++ {
		array = append(array, reward1[i]-reward2[i])
	}
	sort.Ints(array)
	var sum int
	for _, v := range reward2 {
		sum += v
	}
	for i := 1; i <= k; i++ {
		sum += array[len(array)-i]
	}
	return sum
}
