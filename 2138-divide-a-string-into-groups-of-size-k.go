package main

func divideString(s string, k int, fill byte) []string {
	rs := []rune(s)
	n := len(s) % k
	if n != 0 {
		n = k - n
		for i := 0; i < n; i++ {
			rs = append(rs, rune(fill))
		}
	}
	var result []string
	for i := 0; i < len(rs); i += k {
		result = append(result, string(rs[i:i+k]))
	}
	return result
}
