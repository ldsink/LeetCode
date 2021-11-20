package main

func minimumLength(s string) int {
	rs := []rune(s)
	start, end := 0, len(rs)-1
	for ; start < end && rs[start] == rs[end]; start, end = start+1, end-1 {
		i := start + 1
		for ; i < end && rs[start] == rs[i]; i++ {
		}
		start = i - 1
		i = end - 1
		for ; start < i && rs[end] == rs[i]; i-- {
		}
		end = i + 1
	}
	if start < end {
		return end - start + 1
	} else if start == end {
		return 1
	}
	return 0
}
