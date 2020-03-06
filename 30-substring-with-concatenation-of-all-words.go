package main

import "fmt"

func findSubstring(s string, words []string) []int {
	var result []int

	// []
	if len(words) == 0 {
		return result
	}

	// [""]
	length := len(words[0])
	if length == 0 {
		for i := 0; i <= len(s); i++ {
			result = append(result, i)
		}
		return result
	}

	// ["hello"] length=5
	bag := make(map[string]int)
	for _, word := range words {
		bag[word] += 1
	}
	var used, bags []map[string]int
	for i := 0; i < length; i++ {
		m := make(map[string]int)
		used = append(used, m)
		b := make(map[string]int)
		for k, v := range bag {
			b[k] = v
		}
		bags = append(bags, b)
	}
	usedCount := make(map[int]int)

	for i := 0; i < len(s); i++ {
		// out of index
		if i+length > len(s) {
			continue
		}
		word := s[i : i+length]
		offset := i % length
		if val, ok := bags[offset][word]; ok {
			if val <= 0 {
				// instead previous used word
				for start := i - length*usedCount[offset]; ; start += length {
					// pop word
					popWord := s[start : start+length]
					bags[offset][popWord] += 1
					used[offset][popWord] -= 1
					usedCount[offset] -= 1
					if word == popWord {
						break
					}
				}
			}
			// add it to used
			bags[offset][word] -= 1
			used[offset][word] += 1
			usedCount[offset] += 1
		} else {
			b := make(map[string]int)
			for k, v := range bag {
				b[k] = v
			}
			bags[offset] = b
			used[offset] = make(map[string]int)
			usedCount[offset] = 0
		}

		if usedCount[offset] == len(words) {
			result = append(result, i-length*(usedCount[offset]-1))
		}
	}
	return result
}

func main() {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	fmt.Println(findSubstring(s, words))
}
