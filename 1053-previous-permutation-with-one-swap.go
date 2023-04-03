package main

func prevPermOpt1(arr []int) []int {
	success := false
	indexes := make(map[int]int)
	indexes[arr[len(arr)-1]] = len(arr) - 1
	for i := len(arr) - 2; i >= 0 && !success; i-- {
		indexes[arr[i]] = i
		for j := arr[i] - 1; j >= 0; j-- {
			if v, ok := indexes[j]; ok {
				arr[i], arr[v] = arr[v], arr[i]
				success = true
				break
			}
		}
	}
	return arr
}
