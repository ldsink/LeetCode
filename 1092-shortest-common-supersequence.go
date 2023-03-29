package main

func shortestCommonSupersequence(str1 string, str2 string) string {
	cached := make(map[int]int)
	var find func(int, int) int
	find = func(a, b int) int {
		if a == len(str1) {
			return len(str2[b:])
		} else if b == len(str2) {
			return len(str1[a:])
		}
		key := (a << 10) ^ b
		if r, ok := cached[key]; ok {
			return r
		}
		var result int
		if str1[a] == str2[b] {
			result = 1 + find(a+1, b+1)
		} else {
			result = 1 + find(a+1, b)
			if tmp := 1 + find(a, b+1); result > tmp {
				result = tmp
			}
		}
		cached[key] = result
		return result
	}
	minLen := find(0, 0)
	var result []rune
	for a, b := 0, 0; minLen > 0; minLen-- {
		if a == len(str1) {
			result = append(result, []rune(str2[b:])...)
			break
		} else if b == len(str2) {
			result = append(result, []rune(str1[a:])...)
			break
		} else if str1[a] == str2[b] {
			result = append(result, rune(str1[a]))
			a, b = a+1, b+1
		} else if 1+cached[((a+1)<<10)^b] == minLen {
			result = append(result, rune(str1[a]))
			a += 1
		} else {
			result = append(result, rune(str2[b]))
			b += 1
		}
	}
	return string(result)
}
