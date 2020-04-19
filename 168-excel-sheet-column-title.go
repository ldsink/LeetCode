package main

func convertToTitle(n int) string {
	var r []rune
	for ; n != 0; n /= 26 {
		n--
		r = append(r, rune('A'+n%26))
	}
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
