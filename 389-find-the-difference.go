package main

func findTheDifference(s string, t string) byte {
	count := make(map[rune]int)
	for _, r := range []rune(s) {
		count[r] ++
	}
	for _, r := range []rune(t) {
		if count[r] == 0 {
			return byte(r)
		}
		count[r] --
	}
	return byte(0)
}
