package main

func finalString(s string) string {
	var result []rune
	for _, r := range []rune(s) {
		if r == 'i' {
			for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
				result[i], result[j] = result[j], result[i]
			}
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
