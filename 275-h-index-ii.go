package main

func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}
	var l, r int
	for l, r = 0, len(citations)-1; l != r; {
		m := (l + r) >> 1
		if h := len(citations) - m; citations[m] >= h {
			r = m
		} else {
			l = m + 1
		}
	}
	if h := len(citations) - l; citations[l] >= h {
		return h
	}
	return 0
}
