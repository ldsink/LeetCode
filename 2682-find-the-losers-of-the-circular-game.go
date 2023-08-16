package main

func circularGameLosers(n int, k int) []int {
	nums := make([]int, n)
	for idx, multi := 0, 1; ; multi++ {
		nums[idx]++
		if nums[idx] == 2 {
			break
		}
		step := multi * k
		idx = (idx + step) % n
	}
	var answers []int
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			answers = append(answers, i+1)
		}
	}
	return answers
}
