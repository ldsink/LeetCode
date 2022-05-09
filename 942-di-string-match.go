package main

func diStringMatch(s string) []int {
	result := make([]int, len(s)+1)
	for i := 0; i < len(result); i++ {
		result[i] = i
	}
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			continue
		}
		start := i
		for i += 1; i < len(s) && s[i] == 'D'; i++ {
		}
		for s, e := start, i; s < e; s, e = s+1, e-1 {
			result[s], result[e] = result[e], result[s]
		}
	}
	return result
}
