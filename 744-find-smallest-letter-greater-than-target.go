package main

func nextGreatestLetter(letters []byte, target byte) byte {
	chars := make([]bool, 26)
	for _, b := range letters {
		if !chars[b-'a'] {
			chars[b-'a'] = true
		}
	}
	for i := 1; i <= 26; i++ {
		j := (int(target-'a') + i) % 26
		if chars[j] {
			return byte('a' + j)
		}
	}
	return target
}
