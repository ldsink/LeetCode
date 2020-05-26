package main

func firstBadVersion(n int) int {
	if n == 1 {
		return n
	}
	var l, r int
	for l, r = 1, n; l != r; {
		m := (l + r) >> 1
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
