package main

func maxLengthBetweenEqualCharacters(s string) int {
	min := make([]int, 26)
	max := make([]int, 26)
	for i, r := range s {
		idx := r - 'a'
		if min[idx] == 0 && int(r) != int(s[0]) {
			min[idx] = i
		}
		max[idx] = i
	}
	result := -1
	for i := 0; i < 26; i++ {
		if r := max[i] - min[i] - 1; result < r {
			result = r
		}
	}
	return result
}
