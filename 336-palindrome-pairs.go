package main

func palindromePairs(words []string) [][]int {
	var result [][]int
	for idx1 := 0; idx1 < len(words); idx1++ {
		for idx2 := idx1 + 1; idx2 < len(words); idx2++ {
			word := words[idx1] + words[idx2]
			isPalindrome := true
			for i := 0; i < len(word)/2; i++ {
				if word[i] != word[len(word)-1-i] {
					isPalindrome = false
					break
				}
			}
			if isPalindrome {
				result = append(result, []int{idx1, idx2})
			}
			word = words[idx2] + words[idx1]
			isPalindrome = true
			for i := 0; i < len(word)/2; i++ {
				if word[i] != word[len(word)-1-i] {
					isPalindrome = false
					break
				}
			}
			if isPalindrome {
				result = append(result, []int{idx2, idx1})
			}

		}
	}
	return result
}
