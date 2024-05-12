package main

func minDays(n int) int {
	result := make(map[int]int)
	result[1] = 1
	result[2] = 2
	result[3] = 2
	var search func(int) int
	search = func(i int) int {
		if d, ok := result[i]; ok {
			return d
		}
		a := i%2 + search(i/2)
		b := i%3 + search(i/3)
		if a > b {
			a = b
		}
		result[i] = 1 + a
		return result[i]
	}
	return search(n)
}
