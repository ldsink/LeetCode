package main

func minimumSteps(s string) int64 {
	var zeroEnd int
	for ; zeroEnd < len(s) && s[zeroEnd] == '0'; zeroEnd++ {
	}
	var result int64
	for i := zeroEnd + 1; i < len(s); i++ {
		if s[i] == '0' {
			result += int64(i - zeroEnd)
			zeroEnd++
		}
	}
	return result
}
