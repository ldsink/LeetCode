package main

func countNicePairs(nums []int) int {
	const mod = 1e9 + 7
	count := make(map[int]int)
	var result int
	for _, num := range nums {
		var rev int
		for i := num; i != 0; i /= 10 {
			rev *= 10
			rev += i % 10
		}
		result += count[num-rev]
		result %= mod
		count[num-rev]++
	}
	return result
}
