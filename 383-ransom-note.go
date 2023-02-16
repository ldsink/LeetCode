package main

func canConstruct(ransomNote string, magazine string) bool {
	var resource, target [26]int
	for _, r := range magazine {
		resource[r-'a']++
	}
	for _, r := range ransomNote {
		target[r-'a']++
	}
	for i := 0; i < 26; i++ {
		if resource[i] < target[i] {
			return false
		}
	}
	return true
}
