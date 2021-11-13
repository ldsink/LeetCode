package main

func minFlips(target string) int {
	const CLOSE = '0'
	const OPEN = '1'
	status := CLOSE
	var result int
	for i := 0; i < len(target); i++ {
		if rune(target[i]) != status {
			result++
			if status == CLOSE {
				status = OPEN
			} else {
				status = CLOSE
			}
		}
	}
	return result
}
