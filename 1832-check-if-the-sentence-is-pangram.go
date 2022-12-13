package main

func checkIfPangram(sentence string) bool {
	d := make(map[uint8]bool)
	for _, c := range sentence {
		d[uint8(c-'a')] = true
		if len(d) == 26 {
			return true
		}
	}
	return false
}
