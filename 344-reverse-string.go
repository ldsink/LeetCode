package main

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < len(s)/2; i++ {
		s[i], s[j-i] = s[j-i], s[i]
	}
}
