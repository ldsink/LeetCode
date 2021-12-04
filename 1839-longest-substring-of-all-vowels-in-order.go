package main

func longestBeautifulSubstring(word string) int {
	var result int
	vowels := []rune{'e', 'i', 'o', 'u'}
	for i := 0; i < len(word); i++ {
		if word[i] != 'a' {
			continue
		}
		var total int
		for j := i + total; j < len(word); j++ {
			if word[j] == 'a' {
				total++
			} else {
				break
			}
		}
		i += total

		all := true
		for _, vowel := range vowels {
			count := 0
			for j := i + count; j < len(word); j++ {
				if rune(word[j]) == vowel {
					count++
				} else {
					break
				}
			}
			if count == 0 {
				all = false
				break
			}
			total += count
			i += count
		}

		if all && result < total {
			result = total
		}

		i--
	}
	return result
}
