package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var result []rune
	rs := []rune(s)
	circle := (numRows << 1) - 2

	for i := 0; i < numRows; i++ {
		for start := 0; start < len(rs); start += circle {
			left := start + i
			if left >= len(rs) {
				break
			}
			result = append(result, rs[left])

			right := start + circle - i
			if i != 0 && left != right && right < len(rs) {
				result = append(result, rs[right])
			}
		}
	}
	return string(result)
}
