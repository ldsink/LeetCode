package main

func modifyString(s string) string {
	var result []rune
	for i, r := range []rune(s) {
		if r != '?' {
			result = append(result, r)
			continue
		}
		for j := 0; j < 26; j++ {
			if (i == 0 || rune('a'+j) != result[i-1]) && (i+1 == len(s) || s[i+1] == '?' || int(s[i+1]) != 'a'+j) {
				result = append(result, rune('a'+j))
				break
			}
		}
	}
	return string(result)
}
