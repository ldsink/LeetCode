package main

func letterCasePermutation(s string) []string {
	result := []string{s}
	isUpper := func(a rune) bool {
		return 'A' <= a && a <= 'Z'
	}
	isLower := func(a rune) bool {
		return 'a' <= a && a <= 'z'
	}
	mapping := make(map[rune]rune)
	for i := 0; i < 26; i++ {
		mapping[rune('a'+i)] = rune('A' + i)
		mapping[rune('A'+i)] = rune('a' + i)
	}
	for i, r := range []rune(s) {
		if isUpper(r) || isLower(r) {
			l := len(result)
			for j := 0; j < l; j++ {
				rs := []rune(result[j])
				rs[i] = mapping[rs[i]]
				result = append(result, string(rs))
			}
		}
	}
	return result
}
