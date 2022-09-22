package main

func canFormArray(arr []int, pieces [][]int) bool {
	indexes := make(map[int]int)
	for index, piece := range pieces {
		indexes[piece[0]] = index
	}
	var current []int
	for i := 0; i < len(arr); i++ {
		if len(current) == 0 {
			if j, ok := indexes[arr[i]]; !ok {
				return false
			} else {
				current = pieces[j]
			}
		}
		if arr[i] != current[0] {
			return false
		}
		current = current[1:]
	}
	return true
}
