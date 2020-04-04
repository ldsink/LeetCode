package main

func _partition(start int, s string, palindrome [][]int) [][]string {
	var result [][]string
	if start == len(s)-1 {
		return append(result, []string{s[start:]})
	}

	for end := start; end < len(s); end++ {
		// ensure substring is palindrome
		if palindrome[start][end] == 0 {
			palindrome[start][end] = 1
			for i, j := start, end; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					palindrome[start][end] = 2
					break
				}
			}
		}
		if palindrome[start][end] != 1 {
			continue
		}

		if end == len(s)-1 {
			result = append(result, []string{s[start:]})
			continue
		}
		r := _partition(end+1, s, palindrome)
		if len(r) == 0 {
			continue
		}
		for _, ss := range r {
			result = append(result, append([]string{s[start : end+1]}, ss...))
		}
	}

	return result
}

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	var palindrome [][]int
	for i := 0; i < len(s); i++ {
		palindrome = append(palindrome, make([]int, len(s)))
	}
	return _partition(0, s, palindrome)
}
