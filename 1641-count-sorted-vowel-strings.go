package main

func countVowelStrings(n int) int {
	count := [5]int{1, 1, 1, 1, 1}
	for ; n > 1; n-- {
		tmp := count
		for i, j := 4, 0; i >= 0; i-- {
			j += count[i]
			tmp[i] = j
		}
		count = tmp
	}
	var result int
	for _, c := range count {
		result += c
	}
	return result
}
