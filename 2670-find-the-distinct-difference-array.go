package main

func distinctDifferenceArray(nums []int) []int {
	count, prefix := make(map[int]int), make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	var result []int
	for _, n := range nums {
		count[n]--
		if count[n] == 0 {
			delete(count, n)
		}
		prefix[n]++
		result = append(result, len(prefix)-len(count))
	}
	return result
}
