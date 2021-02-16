package main

func reverseVowels(s string) string {
	vowels := make(map[rune]bool)
	vowels['a'] = true
	vowels['e'] = true
	vowels['i'] = true
	vowels['o'] = true
	vowels['u'] = true
	vowels['A'] = true
	vowels['E'] = true
	vowels['I'] = true
	vowels['O'] = true
	vowels['U'] = true

	var normal, reverse []rune
	for _, r := range []rune(s) {
		if _, ok := vowels[r]; ok {
			reverse = append(reverse, r)
		} else {
			normal = append(normal, r)
		}
	}

	for i, j := 0, len(reverse)-1; i < len(reverse)/2; i++ {
		reverse[i], reverse[j-i] = reverse[j-i], reverse[i]
	}

	var result []rune
	var nIndex, rIndex int
	for _, r := range []rune(s) {
		if _, ok := vowels[r]; ok {
			result = append(result, reverse[rIndex])
			rIndex++
		} else {
			result = append(result, normal[nIndex])
			nIndex++
		}
	}

	return string(result)
}
