package main

func buildArray(target []int, n int) []string {
	var result []string
	current := 1
	for _, num := range target {
		for ; current < num; current++ {
			result = append(result, "Push")
			result = append(result, "Pop")
		}
		result = append(result, "Push")
		current++
	}
	return result
}
