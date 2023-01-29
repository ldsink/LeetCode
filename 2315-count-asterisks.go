package main

func countAsterisks(s string) int {
	const Wrap = '|'
	const Asterisk = '*'
	var inWrap bool
	var count int
	for _, r := range s {
		if r == Wrap {
			inWrap = !inWrap
		} else if !inWrap && r == Asterisk {
			count++
		}
	}
	return count
}
